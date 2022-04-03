package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProductCategoryRelDB struct {
	DB
}

func (f *ProductCategoryRelDB) Open() error {
	db, err := sql.Open("mysql",
		"fanland:Password123#@!@tcp(127.0.0.1:3306)/"+f.dbName)
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductCategoryRelDB) GetByRelationships(categoryId uint64) ([]*dao.ProductCategoryRelDO, error) {
	var (
		productId  uint64
		id         uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_id, category_id, create_time, update_time from product_category_rel where category_id = ?", categoryId)

	if err != nil {
		return nil, err
	}

	var relationships []*dao.ProductCategoryRelDO
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &productId, &categoryId, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		err = rows.Err()
		if err != nil {
			return nil, err
		}

		rel := &dao.ProductCategoryRelDO{
			Id:         id,
			ProductId:  productId,
			CategoryId: categoryId,
			CreateTime: createTime,
			UpdateTime: updateTime,
		}
		relationships = append(relationships, rel)
	}
	return relationships, nil
}

func (f *ProductCategoryRelDB) Insert(rel *dao.ProductCategoryRelDO) (err error) {

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

func (f *ProductCategoryDB) Delete(rel *dao.ProductCategoryRelDO) error {

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
