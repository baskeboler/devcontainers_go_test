package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // add this
)

func main() {
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")
	port := 5432
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	r := gin.Default()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	r.GET("/ping", func(c *gin.Context) {
		pingHandler(c, db)
	})
	r.Run()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func pingHandler(c *gin.Context, db *sql.DB) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
