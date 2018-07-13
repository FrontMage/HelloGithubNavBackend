package dao

import (
	"flag"
	"log"
	"time"

	// gorm mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var mysqlURL = flag.String("mysql_url", "root:123456@tcp(localhost:3306)/hello_github", "Specify MySQL URL")
var mysqlConnMaxLifetime = flag.String("mysql_conn_max_lifetime", "+60s", "MySQL connection maximum lifetime, format as +60s +30m etc.")
var mysqlMaxIdleConns = flag.Int("mysql_max_idle_conns", 10, "Maximun idle MySQL connection count")
var mysqlMaxOpenConns = flag.Int("mysql_max_open_conns", 30, "Maximun opening MySQL connection count")

// DB gorm mysql db connection
var DB *gorm.DB

// ErrRecordNotFound 当数据库记录没找到时返回的错误
var ErrRecordNotFound = gorm.ErrRecordNotFound

// init
func init() {
	flag.Parse()
	lifetime, err := time.ParseDuration(*mysqlConnMaxLifetime)
	if err != nil {
		log.Printf("Failed to parse mysql_conn_max_lifetime, err=%s\n", err.Error())
		log.Println("Using defautl lifetime +60s")
	}
	DB, err = gorm.Open("mysql", *mysqlURL)
	if err != nil {
		log.Printf("Failed to connect to mysql with gorm err=%s\n mysqlURL=%s", err.Error(), *mysqlURL)
		return
	}
	db := DB.DB()
	db.SetConnMaxLifetime(lifetime)
	db.SetMaxIdleConns(*mysqlMaxIdleConns)
	db.SetMaxOpenConns(*mysqlMaxOpenConns)
}
