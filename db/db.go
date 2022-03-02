package dao

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type DB struct {
	db     *sql.DB
	dbName string
}

func (f *DB) InitDB(dbName string) {
	f.dbName = dbName
}

func (f *ProductTagDB) Open() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/"+f.dbName)
	f.db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (f *DB) Close() error {
	return f.db.Close()
}
