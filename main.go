package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	db     *sql.DB
)

func init() {
	logger, _ = zap.NewDevelopment()

	dsn := os.Getenv("DATABASE_URL")

	logger.Info(dsn)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(err.Error())

		os.Exit(1)
	}

	if err := database.Ping(); err != nil {
		logger.Error(err.Error())

		os.Exit(1)
	}

	db = database
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Hello Golang!")

	rows, err := db.Query("SELECT 1")
	if err != nil {
		logger.Error(err.Error())
	}

	for rows.Next() {
		var res string
		rows.Scan(&res)
		logger.Info(res)
	}

	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
