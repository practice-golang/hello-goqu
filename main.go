package main // import "hello-goqu"

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlserver"
)

// DB설정
// var (
// 	server   = "127.0.0.1"
// 	port     = 13306
// 	user     = "root"
// 	password = ""
// 	database = "bookshelf"

// 	Dsn *sql.DB
// )

// // initDB - DB파일 생성
// func initDB() (*sql.DB, error) {
// 	var err error
// 	fn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)

// 	Dsn, err = sql.Open("mysql", fn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return Dsn, nil
// }

func main() {
	// Dsn, _ = initDB()

	// db := goqu.New("mysql", Dsn)
	// dl := goqu.Dialect("postgres")
	dl := goqu.Dialect("sqlserver")

	ds := dl.Insert("TableName").Rows(
		goqu.Record{"first_name": "Greg", "last_name": "Farley"},
	)
	insertSQL, args, _ := ds.ToSQL()
	fmt.Println("SQL Log:", insertSQL, args)
}
