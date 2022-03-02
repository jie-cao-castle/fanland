package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProductCategoryDB struct {
	DB
}

func (f *ProductCategoryDB) Open() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/"+f.dbName)
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductCategoryDB) getById(productId int64) (*dao.ProductCategoryDO, error) {
	var (
		name       string
		desc       string
		id         uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, category_name, category_desc ,create_time, update_time from product_category where id = ?", productId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &desc, &createTime, &updateTime)
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

	product := &dao.ProductCategoryDO{
		Id:         id,
		Name:       name,
		Desc:       desc,
		CreateTime: createTime,
		UpdateTime: updateTime,
	}
	return product, nil
}

func (f *ProductCategoryDB) insert(product *dao.ProductCategoryDO) (err error) {

	query := "INSERT INTO product_category (category_name, category_desc, create_time, update_time) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, product.Name, product.Desc)
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

func (f *ProductCategoryDB) update(category *dao.ProductCategoryDO) error {

	query := "UPDATE product_category SET category_name=?, category_desc=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, category.Name, category.Desc, category.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 {
		log.Infof("no update")
	}

	return err
}

func (f *ProductCategoryDB) GetList(limit uint64, offset uint64) ([]*dao.ProductCategoryDO, error) {
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

	rows, err := f.db.Query("select id, category_name, category_desc, create_time, update_time from product_category LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var categories []*dao.ProductCategoryDO
	for rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &nftId, &tags, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		category := &dao.ProductCategoryDO{
			Id:         id,
			Name:       name,
			Desc:       desc,
			CreateTime: createTime,
			UpdateTime: updateTime,
		}

		categories = append(categories, category)
	}
	err = rows.Err()
	if err != nil {
		return categories, err
	}
	return categories, nil
}
