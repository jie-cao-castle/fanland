import (
	"model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fanland/db/product_db"
	log "github.com/sirupsen/logrus"
)

type ProductDB struct {

}

func (f *ProductDB) init() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func (f *ProductDB) getById(id int64)(product *db.ProductDAO, err error) {
	var (
		name string
		desc string
		id int64
		imgUrl string
		nftId int64
		tags string
		createTime time.Time
		updateTime time.Time
	)

	rows, err := db.Query("select id, product_name, desc, imgUrl, nft_id, tag_ids, create_time, update_time from product where id = ?", id)
	
	if err != nil {
		return (nil, err)
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &nftId, &tags, &createTime, &updateTime)
		if err != nil {
			return (nil, err)
		}
	} else {
		return (nil, err)
	}

	err = rows.Err()
	if err != nil {
		return (nil, err)
	}

	product := &db.ProductDAO{
		id: id,
		name: name,
		desc: desc,
		imgUrl: imgUrl,
		nft_id: nftId,
		tag_ids: tags
	}
	return (product, nil)
}

func (f *ProductDB) insert(product *db.ProductDAO)(err error) {

	query := "INSERT INTO product(product_name, desc,imgUrl, nft_id, tag_ids, create_time, update_time) VALUES (?, ?, ? ,?, ? , CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := db.PrepareContext(ctx, query)

    if err != nil {
    	return err
    }
    defer stmt.Close()

	res, err := stmt.ExecContext(ctx, product.name, product.desc, product.imgUrl, product.nft_id, product.tag_ids)  
	if err != nil {  
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()  
	if err != nil {  
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}


func (f *ProductDB) update(product *db.ProductDAO) error {
	insForm, err := db.Prepare("UPDATE Prodect SET name=?, city=? WHERE id=?")

	query := "UPDATE Prodect SET product_name=?, desc=?, tag_ids = ?, update_time = CURRENT_TIMESTAMP WHERE id=?"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, product.name, product.desc, product.tag_ids, id)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return err
}
