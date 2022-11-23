package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Hello Golang!")

	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
