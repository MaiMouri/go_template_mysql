package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerId string
	Name       string
	Age        int
}

func New() *gorm.DB {
	dbconf := "user:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	// データベースのハンドルを取得する
	db, err := gorm.Open("mysql", dbconf)
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}

	db.AutoMigrate(&Customer{})

	return db
}
