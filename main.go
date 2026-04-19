package main

import (
	"concurrencyParser/handlers"
	"concurrencyParser/parser"
	"concurrencyParser/storage"
	"fmt"
	"net/http"
	"os"
)

func main() {
	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=responses sslmode=disable",
		os.Getenv("DB_PASSWORD"),
	)
	db, err := storage.NewDB(connStr)
	if err != nil {
		fmt.Println(err)
	}
	h := handlers.NewHandler(db)
	defer db.Close()
	fmt.Println("Сервер запущен на порту 8080")
	http.HandleFunc("/result", h.HandleGet)
	go parser.ScrapeRia("https://ria.ru", db)
	go parser.ScrapeRbk("https://rbc.ru", db)
	http.ListenAndServe(":8080", nil)
}
