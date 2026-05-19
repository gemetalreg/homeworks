package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {
	http.HandleFunc("/dice", func(w http.ResponseWriter, r *http.Request) {
		num := rand.IntN(6) + 1
		w.Write([]byte(fmt.Sprintf("%d", num)))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
