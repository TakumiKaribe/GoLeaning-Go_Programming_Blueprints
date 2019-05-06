package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/chat/", http.StripPrefix("/chat/", http.FileServer(http.Dir("chat"))))

	// Webサーバーを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
