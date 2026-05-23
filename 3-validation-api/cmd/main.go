package main

import (
	"email/api/verify"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Подключаем метод POST /send
	mux.HandleFunc("POST /send", verify.Send())
	mux.HandleFunc("GET /verify/{hash}", verify.Verify())

	fmt.Println("Сервер запущен на порту :8081")
	http.ListenAndServe(":8081", mux)

}
