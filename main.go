package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"go_todo/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.New()
	fmt.Println(db)
	connectOnly()
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3000")
}

func connectOnly() {
	dbconf := "user:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了")
	}
}
