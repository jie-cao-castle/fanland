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

type ProductDB struct {
	DB
}

func (f *ProductDB) Open() error {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/"+f.dbName+"?parseTime=true")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ProductDB) GetById(productId uint64) (*dao.ProductDO, error) {
	var (
		name        string
		desc        string
		id          uint64
		imgUrl      string
		externalUrl string
		creatorId   uint64
		tags        string
		createTime  time.Time
		updateTime  time.Time
	)

	rows, err := f.db.Query("select id, product_name, product_desc, image_url, external_url, creator_id, tag_ids, create_time, update_time from product where id = ?", productId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &externalUrl, &creatorId, &tags, &createTime, &updateTime)
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
		Id:          id,
		Name:        name,
		Desc:        desc,
		ImgUrl:      imgUrl,
		ExternalUrl: externalUrl,
		CreatorId:   creatorId,
		Tags:        tags,
	}
	return product, nil
}

func (f *ProductDB) GetTitleProduct() (*dao.ProductDO, error) {
	var (
		name        string
		desc        string
		id          uint64
		imgUrl      string
		externalUrl string
		creatorId   uint64
		tags        string
		createTime  time.Time
		updateTime  time.Time
	)

	rows, err := f.db.Query("select id, product_name, product_desc, image_url, external_url, creator_id, tag_ids, create_time, update_time from product LIMIT 1")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &externalUrl, &creatorId, &tags, &createTime, &updateTime)
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
		Id:          id,
		Name:        name,
		Desc:        desc,
		ImgUrl:      imgUrl,
		ExternalUrl: externalUrl,
		CreatorId:   creatorId,
		Tags:        tags,
	}
	return product, nil
}

func (f *ProductDB) GetListByUserId(userId uint64) ([]*dao.ProductDO, error) {
	var (
		name        string
		desc        string
		id          uint64
		imgUrl      string
		externalUrl string
		tags        string
		createTime  time.Time
		updateTime  time.Time
	)

	rows, err := f.db.Query("select id, product_name, product_desc, image_url, external_url, creator_id, tag_ids, create_time, update_time from product where creatorId = ?", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*dao.ProductDO
	for rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &externalUrl, &tags, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		product := &dao.ProductDO{
			Id:          id,
			Name:        name,
			Desc:        desc,
			ImgUrl:      imgUrl,
			ExternalUrl: externalUrl,
			Tags:        tags,
		}

		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (f *ProductDB) Insert(product *dao.ProductDO) (err error) {
	query := "INSERT INTO product(product_name, product_desc, image_url, external_url, creator_id, tag_ids, create_time, update_time) VALUES (?,?,?,?,?,?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, product.Name, product.Desc, product.ImgUrl, product.ExternalUrl, product.CreatorId, product.Tags)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	lastId, _ := res.LastInsertId()
	product.Id = uint64(lastId)

	return nil
}

func (f *ProductDB) Update(product *dao.ProductDO) error {
	query := "UPDATE product SET product_name=?, desc=?, tag_ids = ?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	res, err := stmt.ExecContext(ctx, product.Name, product.Desc, product.Tags, product.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 {
		log.Infof("no update")
	}

	return err
}

func (f *ProductDB) GetList(limit int64, offset int64) ([]*dao.ProductDO, error) {
	var (
		name        string
		desc        string
		id          uint64
		imgUrl      string
		externalUrl string
		tags        string
		creatorId   uint64
		createTime  time.Time
		updateTime  time.Time
	)

	rows, err := f.db.Query("select id, product_name, product_desc, image_url, external_url, creator_id, tag_ids, create_time, update_time from product LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*dao.ProductDO
	for rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &externalUrl, &creatorId, &tags, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		product := &dao.ProductDO{
			Id:          id,
			Name:        name,
			Desc:        desc,
			ImgUrl:      imgUrl,
			ExternalUrl: externalUrl,
			Tags:        tags,
			CreatorId:   creatorId,
			CreateTime:  createTime,
			UpdateTime:  updateTime,
		}

		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (f *ProductDB) GetListByIds(ids []uint64) ([]*dao.ProductDO, error) {
	var (
		name        string
		desc        string
		id          uint64
		imgUrl      string
		externalUrl string
		creatorId   uint64
		tags        string
		createTime  time.Time
		updateTime  time.Time
	)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := f.db.Query("select id, product_name, product_desc, img_url, external_url, creator_id, tag_ids, create_time, update_time from product WHERE id IN (?"+strings.Repeat(",?", len(args)-1)+")", args)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var products []*dao.ProductDO
	for rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &externalUrl, &creatorId, &tags, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		tag := &dao.ProductDO{
			Id:         id,
			Name:       name,
			CreateTime: createTime,
			UpdateTime: updateTime,
		}

		products = append(products, tag)
	}
	err = rows.Err()
	if err != nil {
		return products, err
	}
	return products, nil
}
