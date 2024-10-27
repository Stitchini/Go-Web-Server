package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Capybara struct {
	Name  string
	Color string
}

var capybaras = []Capybara{
	{Name: "Bobby", Color: "Brown"},
	{Name: "Cappy", Color: "Gray"},
	{Name: "Luna", Color: "Biege"},
}

func getAllCapybaras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(capybaras); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
	}
}

func addCapybara(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newCapybara Capybara
	if err := json.NewDecoder(r.Body).Decode(&newCapybara); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	capybaras = append(capybaras, newCapybara)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(capybaras)
}

func main() {
	capybaras = append(capybaras, Capybara{"Globby", "Black"})

	http.HandleFunc("/capybaras", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllCapybaras(w, r)
		case http.MethodPost:
			addCapybara(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my Go Web app")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
