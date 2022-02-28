package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type NftDB struct {
	db     *sql.DB
	dbName string
}

func (f *NftDB) InitDB(dbName string) {
	f.dbName = dbName
}

func (f *NftDB) Open() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/fanland")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	return nil
}

func (f *NftDB) Close() error {
	return f.db.Close()
}

func (f *NftDB) GetById(id uint64) (nftDO *dao.NftDO, err error) {
	var (
		nftId       uint64
		productId   uint64
		productName string
		chainId     uint64
		chainCode   string
		chainName   string
		tokenSymbol string
		tokenName   string
		price       uint64
		priceUnit   uint64
		createTime  time.Time
		updateTime  time.Time
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
		Id:          nftId,
		ProductId:   productId,
		ProductName: productName,
		ChainId:     chainId,
		ChainCode:   chainCode,
		ChainName:   chainName,
		TokenSymbol: tokenSymbol,
		TokenName:   tokenName,
		Price:       price,
		PriceUnit:   priceUnit,
		CreateTime:  createTime,
		UpdateTime:  updateTime,
	}
	return nftObj, nil
}

func (f *NftDB) insert(nft *dao.NftDO) (err error) {

	query := "INSERT INTO nft(product_id, prodect_name, chain_id, chain_code, chain_name, token_symbol, token_name, " +
		"price, price_unit, create_time, update_time) " +
		"VALUES (?, ?, ? ,?, ? ,?, ?, ? ,?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, nft.ProductId, nft.ProductName, nft.ChainId, nft.ChainCode, nft.ChainCode,
		nft.TokenSymbol, nft.TokenName, nft.Price, nft.PriceUnit)
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

func (f *NftDB) update(nft *dao.NftDO) error {
	query := "UPDATE Prodect SET price =?, priceUnit=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, nft.Price, nft.PriceUnit, nft.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowCnt != 1 {
		log.Infof("no update")
	}

	return err
}

func (f *NftDB) getList(limit int64, offset int64) ([]*dao.NftDO, error) {
	var (
		nftId       uint64
		productId   uint64
		productName string
		chainId     uint64
		chainCode   string
		chainName   string
		tokenSymbol string
		tokenName   string
		price       uint64
		priceUnit   uint64
		createTime  time.Time
		updateTime  time.Time
	)
	rows, err := f.db.Query("select id, nft_id, product_id, product_name, chain_id, chain_code, chain_name, "+
		"token_symbol, tokenName, price, priceUnit, create_time, update_time from nft LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var nfts []*dao.NftDO
	for rows.Next() {
		err := rows.Scan(&nftId, &productId, &productName, &chainId, &chainCode, &chainName, &tokenSymbol, &tokenName,
			&price, &priceUnit, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		nftObj := &dao.NftDO{
			Id:          nftId,
			ProductId:   productId,
			ProductName: productName,
			ChainId:     chainId,
			ChainCode:   chainCode,
			ChainName:   chainName,
			TokenSymbol: tokenSymbol,
			TokenName:   tokenName,
			Price:       price,
			PriceUnit:   priceUnit,
			CreateTime:  createTime,
			UpdateTime:  updateTime,
		}
		nfts = append(nfts, nftObj)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return nfts, nil
}
