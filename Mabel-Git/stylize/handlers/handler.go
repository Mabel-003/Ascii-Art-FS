package handlers

import (
	"html/template"
	"log"
	"net/http"
	asciiartgenerator "stylize/asciiArtGenerator"
)

type DataPage struct {
	Result string
	Input  string
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "400 Method Not Allowed", http.StatusBadRequest)
		return
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = temp.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println("Error executing template")
		return
	}

}

func AsciiPage(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
	}()

	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	input := r.FormValue("inputText")
	banner := r.FormValue("banner")

	if input == "" || banner == "" {
		http.Error(w, "400 Bad Request\n \nNeeds an input and banner type..", http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	result := asciiartgenerator.PrintAsciiArt(input, banner)

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := DataPage{
		Result: result,
		Input:  input,
	}

	temp.ExecuteTemplate(w, "index", data)

}
