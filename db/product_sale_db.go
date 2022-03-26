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
		"root:root@tcp(127.0.0.1:3306)/"+f.dbName+"?parseTime=true")
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
		"from_user_id, create_time, update_time) VALUES (?, ?, ? ,?, ? ,?, ?, ? ,?, ?, ?, ?, ?, CURRENT_TIMESTAMP, " +
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
		product.FromUserId)

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
		id            uint64
		productId     uint64
		productName   string
		chainId       uint64
		chainCode     string
		chainName     string
		contractId    uint64
		price         uint64
		priceUnit     uint64
		startTime     time.Time
		endTime       time.Time
		effectiveTime time.Time
		status        int16
		createTime    time.Time
		updateTime    time.Time
		fromUserId    uint64
	)

	rows, err := f.db.Query("select id, product_id, product_name, chain_id, chain_code, chain_name,"+
		"contract_id, price, price_unit, start_time, end_time, effective_time, status, create_time, "+
		"update_time from product_sales WHERE product_id = ? ", queryProductId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var productSales []*dao.ProductSaleDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &productName, &chainId, &chainCode, &chainName, &contractId, &price,
			&priceUnit, &startTime, &endTime, &effectiveTime, &status, &fromUserId, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		productSale := &dao.ProductSaleDO{
			Id:            id,
			ProductId:     productId,
			ProductName:   productName,
			ChainId:       chainId,
			ChainCode:     chainCode,
			ChainName:     chainName,
			ContractId:    contractId,
			Price:         price,
			PriceUnit:     priceUnit,
			StartTime:     startTime,
			EndTime:       endTime,
			EffectiveTime: effectiveTime,
			Status:        status,
			CreateTime:    createTime,
			UpdateTime:    updateTime,
			FromUserId:    fromUserId,
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
		id            uint64
		productId     uint64
		productName   string
		chainId       uint64
		chainCode     string
		chainName     string
		contractId    uint64
		price         uint64
		priceUnit     uint64
		startTime     time.Time
		endTime       time.Time
		effectiveTime time.Time
		status        int16
		createTime    time.Time
		updateTime    time.Time
		fromUserId    uint64
	)

	rows, err := f.db.Query("select id, product_id, product_name, chain_id, chain_code, chain_name,"+
		"contract_id, price, price_unit, start_time, end_time, effective_time, status, create_time, "+
		"update_time from product_sale LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var productSales []*dao.ProductSaleDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &productName, &chainId, &chainCode, &chainName, &contractId, &price,
			&priceUnit, &startTime, &endTime, &effectiveTime, &status, &fromUserId, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		productSale := &dao.ProductSaleDO{
			Id:            id,
			ProductId:     productId,
			ProductName:   productName,
			ChainId:       chainId,
			ChainCode:     chainCode,
			ChainName:     chainName,
			ContractId:    contractId,
			Price:         price,
			PriceUnit:     priceUnit,
			StartTime:     startTime,
			EndTime:       endTime,
			EffectiveTime: effectiveTime,
			Status:        status,
			CreateTime:    createTime,
			UpdateTime:    updateTime,
			FromUserId:    fromUserId,
		}

		productSales = append(productSales, productSale)
	}
	err = rows.Err()
	if err != nil {
		return productSales, err
	}
	return productSales, nil
}
