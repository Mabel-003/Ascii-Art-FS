package main

import (
	"log"
	"net/http"
	"stylize/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/ascii-art", handlers.AsciiPage)
	http.Handle("/styling/", http.StripPrefix("/styling", http.FileServer(http.Dir("./styling"))))

	log.Print("server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)

}
