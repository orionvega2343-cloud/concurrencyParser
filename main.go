package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Сервер запущен на порту 8080")
	http.HandleFunc("/result", HandleGet)
	http.ListenAndServe(":8080", nil)
}
