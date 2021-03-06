package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ChainTokenDB struct {
	DB
}

func (f *ChainTokenDB) init() error {
	db, err := sql.Open("mysql",
		"fanland:Password123#@!@tcp(127.0.0.1:3306)/"+f.dbName)
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *ChainTokenDB) GetById(tokenId uint64) (*dao.ChainTokenDO, error) {
	var (
		id          uint64
		tokenSymbol string
		tokenName   string
		tokenDesc   string
		createTime  time.Time
		updateTime  time.Time
	)

	rows, err := f.db.Query("select id, token_symbol, token_name, token_desc, create_time, update_time from chain_token where id = ?", tokenId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id, &tokenSymbol, &tokenName, &tokenDesc, &createTime, &updateTime)
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
	token := &dao.ChainTokenDO{
		Id:          id,
		TokenName:   tokenName,
		TokenSymbol: tokenSymbol,
		TokenDesc:   tokenDesc,
		CreateTime:  createTime,
		UpdateTime:  updateTime,
	}

	return token, nil
}

func (f *ChainTokenDB) insert(token *dao.ChainTokenDO) (err error) {

	query := "INSERT INTO chain_token(token_symbol, token_name, token_desc, create_time, update_time) VALUES (?, ?, ? , CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, token.TokenSymbol, token.TokenName, token.TokenDesc)
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

func (f *ChainTokenDB) update(token *dao.ChainTokenDO) error {

	query := "UPDATE chain_token SET token_symbol=?, token_name=?, token_desc=?, tag_ids = ?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return err
	}
	res, err := stmt.ExecContext(ctx, token.TokenSymbol, token.TokenName, token.TokenDesc, token.Id)
	if err != nil {
		log.Error(err)
		return err
	}

	rowCnt, err := res.RowsAffected()
	if rowCnt != 1 {
		log.Info("no update")
		return err
	}

	return err
}

func (f *ChainTokenDB) getList(limit int64, offset int64) ([]*dao.ChainTokenDO, error) {
	var (
		id          uint64
		tokenSymbol string
		tokenName   string
		tokenDesc   string
		createTime  time.Time
		updateTime  time.Time
	)

	rows, err := f.db.Query("select id, tokenSymbol, tokenName, tokenDesc, create_time, update_time from chain_token LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tokens []*dao.ChainTokenDO
	for rows.Next() {
		err := rows.Scan(&id, &tokenSymbol, &tokenName, &tokenDesc, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		token := &dao.ChainTokenDO{
			Id:          id,
			TokenName:   tokenName,
			TokenSymbol: tokenSymbol,
			TokenDesc:   tokenDesc,
			CreateTime:  createTime,
			UpdateTime:  updateTime,
		}

		tokens = append(tokens, token)
	}
	err = rows.Err()
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}
