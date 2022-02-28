package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProductOrderDB struct {
	db *sql.DB
	DB
}

func (f *ProductOrderDB) init() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/fanland")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductOrderDB) getById(orderId uint64) (*dao.ProductOrderDO, error) {
	var (
		id         uint64
		productId  uint64
		offerId    uint64
		nftId      uint64
		price      uint64
		nftUnit    uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_name, product_desc, imgUrl, nft_id, tag_ids, create_time, update_time from product where id = ?", orderId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &productId, &offerId, &nftId, &nftId, &price, &nftUnit, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	productOrder := &dao.ProductOrderDO{
		Id:         id,
		ProductId:  productId,
		OfferId:    offerId,
		NftId:      nftId,
		Price:      price,
		NftUnit:    nftUnit,
		CreateTime: createTime,
		UpdateTime: updateTime,
	}
	return productOrder, nil
}

func (f *ProductOrderDB) insert(productOrder *dao.ProductOrderDO) (err error) {

	query := "INSERT INTO product(product_name, desc,imgUrl, nft_id, tag_ids, create_time, update_time) VALUES (?, ?, ? ,?, ? , CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, productOrder.ProductId, productOrder.OfferId, productOrder.NftUnit, productOrder.Price, productOrder.NftUnit)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *ProductOrderDB) update(productOrder *dao.ProductOrderDO) error {

	query := "UPDATE Prodect SET price =?, nft_unit=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, productOrder.Price, productOrder.NftUnit, productOrder.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 {
		log.Infof("no update")
	}

	return err
}

func (f *ProductOrderDB) getList(limit int64, offset int64) ([]*dao.ProductOrderDO, error) {
	var (
		id         uint64
		productId  uint64
		offerId    uint64
		nftId      uint64
		price      uint64
		nftUnit    uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_name, desc, imgUrl, nft_id, tag_ids, create_time, update_time from product LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var productOrders []*dao.ProductOrderDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &offerId, &nftId, &nftId, &price, &nftUnit, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		productOrder := &dao.ProductOrderDO{
			Id:         id,
			ProductId:  productId,
			OfferId:    offerId,
			NftId:      nftId,
			Price:      price,
			NftUnit:    nftUnit,
			CreateTime: createTime,
			UpdateTime: updateTime,
		}

		productOrders = append(productOrders, productOrder)
	}
	err = rows.Err()
	if err != nil {
		return productOrders, err
	}
	return productOrders, nil
}
