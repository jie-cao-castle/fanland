package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	log "github.com/sirupsen/logrus"
	"time"
)

type NftOrderDB struct {
	DB
}

func (f *NftOrderDB) Open() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/"+f.dbName)
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *NftOrderDB) Insert(nftOrder *dao.NftOrderDO) (err error) {

	query := "INSERT INTO nft_order(product_id, chain_id, chain_code, nft_key, price, price_unit, amount, order_status " +
		"transaction_hash, create_time, update_time) " +
		"VALUES (?, ?, ? ,?, ? ,?, ?, ? ,?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, nftOrder.ProductId, nftOrder.ChainId, nftOrder.ChainCode,
		nftOrder.NftKey, nftOrder.Price, nftOrder.PriceUnit, nftOrder.Amount, nftOrder.Status, nftOrder.TransactionHash)
	if err != nil {
		log.Errorf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Errorf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *NftOrderDB) GetListByProductId(queryProductId uint64) ([]*dao.NftOrderDO, error) {
	var (
		id              uint64
		productId       uint64
		nftKey          string
		price           uint64
		priceUnit       uint64
		amount          uint64
		status          int8
		chainId         uint64
		chainCode       string
		transactionHash string
		createTime      time.Time
		updateTime      time.Time
	)

	rows, err := f.db.Query("select id, product_id, nft_key, price, price_unit, amount, status, chain_id, "+
		"chain_code, transaction_hash, create_time, update_time from product_category WHENEVER product_id = ? ", queryProductId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var nftOrders []*dao.NftOrderDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &nftKey, &price, &priceUnit, &amount, &status, &chainId, &chainCode, &transactionHash, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		nftOrder := &dao.NftOrderDO{
			Id:              id,
			ProductId:       productId,
			NftKey:          nftKey,
			Price:           price,
			PriceUnit:       priceUnit,
			Amount:          amount,
			Status:          status,
			ChainId:         chainId,
			ChainCode:       chainCode,
			TransactionHash: transactionHash,
			CreateTime:      createTime,
			UpdateTime:      updateTime,
		}

		nftOrders = append(nftOrders, nftOrder)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return nftOrders, nil
}

func (f *NftOrderDB) Update(nftOrder *dao.NftOrderDO) (err error) {

	query := "UPDATE nft_contract SET status=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, nftOrder.Status, nftOrder.Id)
	if err != nil {
		log.Errorf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Errorf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}
