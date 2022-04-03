package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProductTagRelDB struct {
	DB
}

func (f *ProductTagRelDB) Open() error {
	db, err := sql.Open("mysql",
		"fanland:Password123#@!@tcp(127.0.0.1:3306)/"+f.dbName)
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductTagRelDB) GetListByTagId(tagId uint64) ([]*dao.ProductTagRelDO, error) {
	var (
		id         uint64
		productId  uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_id, tag_id, create_time, update_time from product_tag_rel where tag_id = ?", tagId)

	if err != nil {
		return nil, err
	}

	var relationships []*dao.ProductTagRelDO
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &productId, &tagId, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		err = rows.Err()
		if err != nil {
			return nil, err
		}

		rel := &dao.ProductTagRelDO{
			Id:         id,
			ProductId:  productId,
			TagId:      tagId,
			CreateTime: createTime,
			UpdateTime: updateTime,
		}
		relationships = append(relationships, rel)
	}
	return relationships, nil
}

func (f *ProductTagRelDB) GetListByProductId(productId uint64) ([]*dao.ProductTagRelDO, error) {
	var (
		tagId      uint64
		id         uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_id, tag_id, create_time, update_time from product_category_rel where product_id = ?", productId)

	if err != nil {
		return nil, err
	}

	var relationships []*dao.ProductTagRelDO
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &productId, &tagId, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		err = rows.Err()
		if err != nil {
			return nil, err
		}

		rel := &dao.ProductTagRelDO{
			Id:         id,
			ProductId:  productId,
			TagId:      tagId,
			CreateTime: createTime,
			UpdateTime: updateTime,
		}
		relationships = append(relationships, rel)
	}
	return relationships, nil
}

func (f *ProductTagRelDB) Insert(rel *dao.ProductCategoryRelDO) (err error) {

	query := "INSERT INTO product_category_rel (product_id, category_id, create_time, update_time) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, rel.ProductId, rel.CategoryId)
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

func (f *ProductTagRelDB) Delete(rel *dao.ProductCategoryRelDO) error {

	query := "DELETE FROM product_category_rel WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, rel.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 {
		log.Infof("no update")
	}

	return err
}
