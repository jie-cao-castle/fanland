package dao

type DB struct {
	dbName string
}

func (f *DB) InitDB(dbName string) {
	f.dbName = dbName
}
