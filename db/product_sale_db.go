package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProductSaleDB struct {
	DB
}

func (f *ProductSaleDB) Open() error {
	db, err := sql.Open("mysql",
		"fanland:Password123#@!@tcp(127.0.0.1:3306)/"+f.dbName+"?parseTime=true")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductSaleDB) Insert(product *dao.ProductSaleDO) (err error) {

	query := "INSERT INTO product_sale(product_id, product_name, chain_id, chain_code, chain_name, contract_id, " +
		"price, price_unit, start_time, end_time, effective_time, sale_status, " +
		"from_user_id, token_id, transaction_hash, create_time, update_time) VALUES (?, ?, ? ,?, ? ,?, ?, ? ,?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, " +
		"CURRENT_TIMESTAMP)"

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx,
		product.ProductId,
		product.ProductName,
		product.ChainId,
		product.ChainCode,
		product.ChainName,
		product.ContractId,
		product.Price,
		product.PriceUnit,
		product.StartTime,
		product.EndTime,
		product.EffectiveTime,
		product.Status,
		product.FromUserId,
		product.TokenId,
		product.TransactionHash)

	if err != nil {
		log.Printf("Error %s when inserting row into product_sale table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *ProductSaleDB) Update(product *dao.ProductSaleDO) (err error) {

	query := "UPDATE product_sale SET sale_status=?, update_time = CURRENT_TIMESTAMP WHERE id=?"

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, product.Status, product.Id)

	if err != nil {
		log.Printf("Error %s when inserting row into product_sale table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *ProductSaleDB) GetListByProductId(queryProductId uint64) ([]*dao.ProductSaleDO, error) {
	var (
		id              uint64
		productId       uint64
		productName     string
		chainId         uint64
		chainCode       string
		chainName       string
		contractId      uint64
		price           uint64
		priceUnit       uint64
		startTime       time.Time
		endTime         time.Time
		effectiveTime   time.Time
		status          int16
		createTime      time.Time
		updateTime      time.Time
		fromUserId      uint64
		fromUserName    string
		tokenId         string
		transactionHash string
	)

	rows, err := f.db.Query("select s.id, s.product_id, s.product_name, s.chain_id, s.chain_code, s.chain_name,"+
		"s.contract_id, s.price, s.price_unit, s.start_time, s.end_time, s.effective_time, s.sale_status, s.from_user_id, u.user_name, s.token_id, s.transaction_hash, s.create_time, "+
		"s.update_time from product_sale s INNER JOIN fanland_user u ON s.from_user_id = u.id WHERE s.product_id = ? ORDER BY s.update_time DESC", queryProductId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var productSales []*dao.ProductSaleDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &productName, &chainId, &chainCode, &chainName, &contractId, &price,
			&priceUnit, &startTime, &endTime, &effectiveTime, &status, &fromUserId, &fromUserName, &tokenId, &transactionHash, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		productSale := &dao.ProductSaleDO{
			Id:              id,
			ProductId:       productId,
			ProductName:     productName,
			ChainId:         chainId,
			ChainCode:       chainCode,
			ChainName:       chainName,
			ContractId:      contractId,
			Price:           price,
			PriceUnit:       priceUnit,
			StartTime:       startTime,
			EndTime:         endTime,
			EffectiveTime:   effectiveTime,
			Status:          status,
			CreateTime:      createTime,
			UpdateTime:      updateTime,
			FromUserId:      fromUserId,
			FromUserName:    fromUserName,
			TokenId:         tokenId,
			TransactionHash: transactionHash,
		}

		productSales = append(productSales, productSale)
	}
	err = rows.Err()
	if err != nil {
		return productSales, err
	}
	return productSales, nil
}

func (f *ProductSaleDB) GetList(limit int64, offset int64) ([]*dao.ProductSaleDO, error) {
	var (
		id              uint64
		productId       uint64
		productName     string
		chainId         uint64
		chainCode       string
		chainName       string
		contractId      uint64
		price           uint64
		priceUnit       uint64
		startTime       time.Time
		endTime         time.Time
		effectiveTime   time.Time
		status          int16
		createTime      time.Time
		updateTime      time.Time
		fromUserId      uint64
		tokenId         string
		transactionHash string
	)

	rows, err := f.db.Query("select id, product_id, product_name, chain_id, chain_code, chain_name,"+
		"contract_id, price, price_unit, start_time, end_time, effective_time, sale_status, from_user_id, token_id, transaction_hash, create_time, "+
		"update_time from product_sale LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var productSales []*dao.ProductSaleDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &productName, &chainId, &chainCode, &chainName, &contractId, &price,
			&priceUnit, &startTime, &endTime, &effectiveTime, &status, &fromUserId, &tokenId, &transactionHash, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		productSale := &dao.ProductSaleDO{
			Id:              id,
			ProductId:       productId,
			ProductName:     productName,
			ChainId:         chainId,
			ChainCode:       chainCode,
			ChainName:       chainName,
			ContractId:      contractId,
			Price:           price,
			PriceUnit:       priceUnit,
			StartTime:       startTime,
			EndTime:         endTime,
			EffectiveTime:   effectiveTime,
			Status:          status,
			CreateTime:      createTime,
			UpdateTime:      updateTime,
			FromUserId:      fromUserId,
			TokenId:         tokenId,
			TransactionHash: transactionHash,
		}

		productSales = append(productSales, productSale)
	}
	err = rows.Err()
	if err != nil {
		return productSales, err
	}
	return productSales, nil
}
