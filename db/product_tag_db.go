package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type ProductTagDB struct {
	db *sql.DB
	DB
}

func (f *ProductTagDB) Init() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/fanland")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductTagDB) Close() error {
	return f.db.Close()
}

func (f *ProductTagDB) GetById(productId int64) (*dao.ProductTagDO, error) {
	var (
		name       string
		id         uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, category_name, create_time, update_time from product_category where id = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &createTime, &updateTime)
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

	product := &dao.ProductTagDO{
		Id:   id,
		Name: name,
	}
	return product, nil
}

func (f *ProductTagDB) insert(tag *dao.ProductTagDO) (err error) {

	query := "INSERT INTO product_category (category_name, category_desc, create_time, update_time) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, tag.Name)
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

func (f *ProductTagDB) update(category *dao.ProductTagDO) error {

	query := "UPDATE product_category SET category_name=?, category_desc=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, category.Name)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 {
		log.Infof("no update")
	}

	return err
}

func (f *ProductTagDB) getList(limit int64, offset int64) ([]*dao.ProductTagDO, error) {
	var (
		name       string
		id         uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, category_name,category_desc, create_time, update_time from product_category LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tags []*dao.ProductTagDO
	for rows.Next() {
		err := rows.Scan(&id, &name, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		tag := &dao.ProductTagDO{
			Id:   id,
			Name: name,
		}

		tags = append(tags, tag)
	}
	err = rows.Err()
	if err != nil {
		return tags, err
	}
	return tags, nil
}

func (f *ProductTagDB) GetListByIds(ids []uint64) ([]*dao.ProductTagDO, error) {
	var (
		name       string
		id         uint64
		createTime time.Time
		updateTime time.Time
	)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := f.db.Query("select id, category_name,category_desc, create_time, update_time from product_category WHERE id IN (?"+strings.Repeat(",?", len(args)-1)+")", args)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var tags []*dao.ProductTagDO
	for rows.Next() {
		err := rows.Scan(&id, &name, &tags, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		tag := &dao.ProductTagDO{
			Id:   id,
			Name: name,
		}

		tags = append(tags, tag)
	}
	err = rows.Err()
	if err != nil {
		return tags, err
	}
	return tags, nil
}
