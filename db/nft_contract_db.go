package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type NftContractDB struct {
	DB
}

func (f *NftContractDB) Open() error {
	db, err := sql.Open("mysql",
		"fanland:Password123#@!@tcp(127.0.0.1:3306)/"+f.dbName+"?parseTime=true")
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *NftContractDB) Insert(nftContract *dao.NftContractDO) (err error) {

	query := "INSERT INTO nft_contract(product_id, chain_id, chain_code, token_symbol, token_name, " +
		"contract_address, contract_status, create_time, update_time) " +
		"VALUES (?, ?, ? ,?, ? ,?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, nftContract.ProductId, nftContract.ChainId, nftContract.ChainCode,
		nftContract.TokenSymbol, nftContract.TokenName, nftContract.ContractAddress, nftContract.Status)
	if err != nil {
		log.Errorf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Errorf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *NftContractDB) Update(nftContract *dao.NftContractDO) (err error) {

	query := "UPDATE nft_contract SET contract_status=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, nftContract.Status, nftContract.Id)
	if err != nil {
		log.Errorf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Errorf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *NftContractDB) UpdateToken(nftContract *dao.NftContractDO) (err error) {

	query := "UPDATE nft_contract SET next_token_id=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, nftContract.NextTokenId, nftContract.Id)
	if err != nil {
		log.Errorf("Error %s when inserting row into products table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Errorf("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *NftContractDB) GetListByProductId(queryProductId uint64) ([]*dao.NftContractDO, error) {
	var (
		id              uint64
		productId       uint64
		chainId         uint64
		chainCode       string
		contractAddress string
		status          int8

		tokenSymbol string
		tokenName   string
		tokenAmount uint64
		nextTokenId uint64

		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_id, chain_id, chain_code, "+
		"contract_address, contract_status, token_symbol, token_name, token_amount, next_token_id, create_time, "+
		"update_time from nft_contract WHERE product_id = ? ", queryProductId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var contractDOs []*dao.NftContractDO
	for rows.Next() {
		err := rows.Scan(&id, &productId, &chainId, &chainCode, &contractAddress, &status,
			&tokenSymbol, &tokenName, &tokenAmount, &nextTokenId, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		contractDO := &dao.NftContractDO{
			Id:              id,
			ProductId:       productId,
			ChainId:         chainId,
			ChainCode:       chainCode,
			ContractAddress: contractAddress,
			Status:          status,
			TokenSymbol:     tokenSymbol,
			TokenName:       tokenName,
			CreateTime:      createTime,
			UpdateTime:      updateTime,
		}

		contractDOs = append(contractDOs, contractDO)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return contractDOs, nil
}

func (f *NftContractDB) GetByChainId(queryProductId uint64, queryChainId uint64) (*dao.NftContractDO, error) {
	var (
		id              uint64
		productId       uint64
		chainId         uint64
		chainCode       string
		contractAddress string
		status          int8

		tokenSymbol string
		tokenName   string
		tokenAmount uint64
		nextTokenId uint64

		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, product_id, chain_id, chain_code, "+
		"contract_address, contract_status, token_symbol, token_name, token_amount, next_token_id, create_time, "+
		"update_time from nft_contract WHERE product_id = ? AND chain_id = ? LIMIT 1", queryProductId, queryChainId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var contractDO *dao.NftContractDO
	if rows.Next() {
		err := rows.Scan(&id, &productId, &chainId, &chainCode, &contractAddress, &status,
			&tokenSymbol, &tokenName, &tokenAmount, &nextTokenId, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		contractDO = &dao.NftContractDO{
			Id:              id,
			ProductId:       productId,
			ChainId:         chainId,
			ChainCode:       chainCode,
			ContractAddress: contractAddress,
			Status:          status,
			TokenSymbol:     tokenSymbol,
			TokenName:       tokenName,
			TokenAmount:     tokenAmount,
			NextTokenId:     nextTokenId,
			CreateTime:      createTime,
			UpdateTime:      updateTime,
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return contractDO, nil
}
