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
		"root:root@tcp(127.0.0.1:3306)/"+f.dbName+"?parseTime=true")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *NftOrderDB) Insert(nftOrder *dao.NftOrderDO) (err error) {

	query := "INSERT INTO nft_order(product_id, chain_id, chain_code, nft_key, price, price_unit, amount, order_status, " +
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
		toUserId        uint64
		toUserName      string
		createTime      time.Time
		updateTime      time.Time
	)

	rows, err := f.db.Query("select o.id, o.product_id, o.nft_key, o.price, o.price_unit, o.amount, o.order_status, o.chain_id, "+
		"o.chain_code, o.transaction_hash, o.create_time, o.update_time, u.id as to_user_id, u.user_name as to_user_name from nft_order o INNER JOIN fanland_user u ON u.id = o.to_user_id WHERE o.product_id = ? ", queryProductId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var nftOrders []*dao.NftOrderDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &nftKey, &price, &priceUnit, &amount, &status, &chainId, &chainCode,
			&transactionHash, &createTime, &updateTime, &toUserId, &toUserName)
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
			ToUserId:        toUserId,
			ToUserName:      toUserName,
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

	query := "UPDATE nft_order SET order_status = ?, update_time = CURRENT_TIMESTAMP WHERE id=?"
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
