import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fanland/db/product_db"
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

func (f *ProductDB) getById(id int64)(*Product) {
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (f *ProductDB) update(product *ProductDB) error {
	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Dolly")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	    // 插入数据
		stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
		checkErr(err)
	
		res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
		checkErr(err)
	
		id, err := res.LastInsertId()
		checkErr(err)
	
		fmt.Println(id)
		// 更新数据
		stmt, err = db.Prepare("update userinfo set username=? where uid=?")
		checkErr(err)
	
		res, err = stmt.Exec("astaxieupdate", id)
		checkErr(err)
	
		affect, err := res.RowsAffected()
		checkErr(err)
	
}
