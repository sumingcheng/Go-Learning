package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 你的主要逻辑
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	finalHandler := Chain(mainHandler, LoggingMiddleware, AuthMiddleware)

	http.HandleFunc("/", finalHandler)
	log.Println("Starting server on :9999")
	http.ListenAndServe(":8080", nil)
}
