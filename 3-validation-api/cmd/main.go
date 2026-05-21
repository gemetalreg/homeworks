package main

import (
	"email/api/verify"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Подключаем метод POST /send
	mux.HandleFunc("POST /send", verify.NewHandler())
	mux.HandleFunc("GET /verify/{hash}", verify.NewHandler())

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", mux)

}
