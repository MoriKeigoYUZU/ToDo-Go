package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //使わない
)

var (
	DB = &sql.DB{}
)

// connectLocalSQL localのmysqlのコネクション作成
func ConnectLocalSQL() {
	log.Println("connectDB: local")
	//環境変数->コンピュータ上に設定されている設定用変数
	dbuser := os.Getenv("MYSQL_USER")
	if dbuser == "" {
		dbuser = "root"
	}
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	if dbpassword == "" {
		dbpassword = "password"
	}
	dbhost := os.Getenv("MYSQL_HOST")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbname := os.Getenv("MYSQL_DATABASE")
	if dbname == "" {
		dbname = "dbname"
	}
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbname)
	var err error
	DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
}
