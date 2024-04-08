package main

import (
	"asciiArt/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiHandler)
	http.HandleFunc("/export", handlers.ExportHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
