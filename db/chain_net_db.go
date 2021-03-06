package dao

import (
	"context"
	"database/sql"
	"fanland/db/dao"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

type ChainNetDB struct {
	DB
}

func (f *ChainNetDB) Init() error {
	db, err := sql.Open("mysql",
		"fanland:Password123#@!@tcp(127.0.0.1:3306)/"+f.dbName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	f.db = db
	return nil
}

func (f *ChainNetDB) GetByChainId(queryChainId uint64) (chainNet *dao.ChainNetDO, err error) {
	var (
		dbId       uint64
		chainId    uint64
		chainCode  string
		chainName  string
		createTime time.Time
		updateTime time.Time
	)

	rows, err := f.db.Query("select id, chain_code, chain_id, chain_name, create_time, update_time from chain_net where chain_id = ?", queryChainId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&dbId, &chainId, &chainCode, &chainName, &createTime, &updateTime)
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

	chainNetObj := &dao.ChainNetDO{
		Id:         dbId,
		ChainId:    chainId,
		ChainCode:  chainCode,
		ChainName:  chainName,
		CreateTime: createTime,
		UpdateTime: updateTime,
	}
	return chainNetObj, nil
}

func (f *ChainNetDB) insert(chain *dao.ChainNetDO) (err error) {

	query := "INSERT INTO chain_net(chain_code, chain_id, chain_name, create_time, update_time) " +
		"VALUES (?, ?, ? CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, chain.ChainCode, chain.ChainId, chain.ChainName)
	if err != nil {
		log.Infof("Error %s when inserting row into chain net table", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Infof("Error %s when finding rows affected", err)
		return err
	}

	return nil
}

func (f *ChainNetDB) update(chainNet *dao.ChainNetDO) error {
	query := "UPDATE chain_net SET chain_code =?, chain_id =?, chain_name=?, update_time = CURRENT_TIMESTAMP WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return err
	}
	res, err := stmt.ExecContext(ctx, chainNet.ChainCode, chainNet.ChainId, chainNet.ChainName, chainNet.Id)
	if err != nil {
		log.Error(err)
		return err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}

	if rowCnt != 1 {
		log.Info("no update")
		return nil
	}

	return err
}

func (f *ChainNetDB) getList(limit int64, offset int64) ([]*dao.ChainNetDO, error) {
	var (
		chainId    uint64
		chainCode  string
		chainName  string
		createTime time.Time
		updateTime time.Time
	)
	rows, err := f.db.Query("select id, chain_code, chain_name, create_time, update_time from chain_net LIMIT ? OFFSET ? ", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var chainNets []*dao.ChainNetDO
	for rows.Next() {
		err := rows.Scan(&chainId, &chainCode, &chainName, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		chainNetObj := &dao.ChainNetDO{
			Id:         chainId,
			ChainCode:  chainCode,
			ChainName:  chainName,
			CreateTime: createTime,
			UpdateTime: updateTime,
		}
		chainNets = append(chainNets, chainNetObj)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return chainNets, nil
}
