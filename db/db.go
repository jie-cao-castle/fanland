package dao

import "database/sql"

type DB struct {
	db     *sql.DB
	dbName string
}

func (f *DB) InitDB(dbName string) {
	f.dbName = dbName
}

func (f *DB) Close() error {
	return f.db.Close()
}
