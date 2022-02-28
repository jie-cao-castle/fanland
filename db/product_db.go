package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProductDB struct {
	db *sql.DB
	DB
}

func (f *ProductDB) Open() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/fanland")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductDB) Close() error {
	return f.db.Close()
}

func (f *ProductDB) GetById(productId uint64) (*dao.ProductDO, error) {
	var (
		name       string
		desc       string
		id         uint64
		imgUrl     string
		nftId      uint64
		tags       string
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_name, product_desc, imgUrl, nft_id, tag_ids, create_time, update_time from product where id = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer f.db.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &nftId, &tags, &createTime, &updateTime)
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

	product := &dao.ProductDO{
		Id:     id,
		Name:   name,
		Desc:   desc,
		ImgUrl: imgUrl,
		NftId:  nftId,
		Tags:   tags,
	}
	return product, nil
}

func (f *ProductDB) insert(product *dao.ProductDO) (err error) {

	query := "INSERT INTO product(product_name, desc,imgUrl, nft_id, tag_ids, create_time, update_time) VALUES (?, ?, ? ,?, ? , CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, product.Name, product.Desc, product.ImgUrl, product.NftId, product.Tags)
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

func (f *ProductDB) update(product *dao.ProductDO) error {

	query := "UPDATE Prodect SET product_name=?, desc=?, tag_ids = ?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, product.Name, product.Desc, product.Tags, product.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 {
		log.Infof("no update")
	}

	return err
}

func (f *ProductDB) getList(limit int64, offset int64) ([]*dao.ProductDO, error) {
	var (
		name       string
		desc       string
		id         uint64
		imgUrl     string
		nftId      uint64
		tags       string
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_name, desc, imgUrl, nft_id, tag_ids, create_time, update_time from product LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*dao.ProductDO
	for rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &nftId, &tags, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		product := &dao.ProductDO{
			Id:     id,
			Name:   name,
			Desc:   desc,
			ImgUrl: imgUrl,
			NftId:  nftId,
			Tags:   tags,
		}

		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return products, err
	}
	return products, nil
}
