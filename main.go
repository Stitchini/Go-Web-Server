package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Capybara struct {
	id    int
	Name  string
	Color string
}

var capybaras = []Capybara{
	{id: 1, Name: "Bobby", Color: "Brown"},
	{id: 2, Name: "Cappy", Color: "Gray"},
	{id: 3, Name: "Luna", Color: "Biege"},
}
var nextID = 4

func getAllCapybaras(w http.ResponseWriter, _ *http.Request) {
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
	newCapybara.id = nextID
	nextID++

	capybaras = append(capybaras, newCapybara)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(capybaras)
}

func main() {

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
