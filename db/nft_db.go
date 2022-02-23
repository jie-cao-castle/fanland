package dao
import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
	"fanland/db/dao"
)

type NftDB struct {
	db *sql.DB;
}

func (f *NftDB) init() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/fanland")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	return nil
}

func (f *NftDB) GetById(id uint64)(nftDO *dao.NftDO, err error) {
	var (
		nftId uint64
		productId uint64
		productName string
		chainId uint64
		chainCode string
		chainName string
		tokenSymbol string
		tokenName string
		price uint64
		priceUnit uint64
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_id, prodect_name, chain_id, chain_code, chain_name, token_symbol, token_name, price, price_unit, create_time, update_time from nft where id = ?", id)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&nftId, &productId, &productName, &chainId, &chainCode, &chainName, &tokenSymbol, &tokenName, &price, &priceUnit, &createTime, &updateTime)
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

	nftObj := &dao.NftDO{
		Id: nftId,
		ProductId: productId,
		productName string
		chainId int64
		chainCode string
		chainName string
		tokenSymbol string
		tokenName string
		price int64
		priceUnit int64
		createTime time.Time
		updateTime time.Time
	}
	return nftObj, nil
}

func (f *NftDB) insert(nft *dao.NftDO)(err error) {

	query := "INSERT INTO product(product_id, prodect_name, chain_id, chain_code, chain_name, token_symbol, token_name, price, price_unit, create_time, update_time) VALUES (?, ?, ? ,?, ? , CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := f.db.PrepareContext(ctx, query)

    if err != nil {
    	return err
    }
    defer stmt.Close()

	res, err := stmt.ExecContext(ctx, nft.name, nft.desc, nft.imgUrl, nft.nft_id, nft.tag_ids)
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


func (f *NftDB) update(product *db.ProductDAO) error {
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


func (f *NftDB) getList(limit int64, offset int64)(products []*db.ProductDAO, err error) {
	var (
		id int64
		productId int64
		productName string
		chainId int64
		chainCode string
		chainName string
		tokenSymbol string
		tokenName string
		price int64
		priceUnit int64
		createTime time.Time
		updateTime time.Time
	)
	rows, err := db.Query("select id, product_name, desc, imgUrl, nft_id, tag_ids, create_time, update_time from product LIMIT ? OFFSET ? ", limit, offset)
	
	if err != nil {
		return (nil, err)
	}

	defer rows.Close()
	nfts = []
	for rows.Next() {
		err := rows.Scan(&id, &name, &desc, &imgUrl, &nftId, &tags, &createTime, &updateTime)
		if err != nil {
			return (nil, err)
		}

		nft := &db.NftDAO{
			id : id
			productId : productId
			productName : productName
			chainId : chainId
			chainCode : chainCode
			chainName : chainName
			tokenSymbol : tokenSymbol
			tokenName : tokenName
			price : price
			priceUnit : priceUnit
			createTime : createTime
			updateTime :updateTime
		}
	}

	err = rows.Err()
	if err != nil {
		return ([], err)
	}
	return (nfts, nil)
}