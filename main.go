package main

import (
	"asciiArt/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiHandler)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
