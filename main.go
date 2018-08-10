package main

import (
	"../github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
)

type Books struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author string `json:"author"`

}

var books []Books

// Get All Books
func getBooks(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r * http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)

    for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Books{})
}

// Create A New Book
func createBook(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type","application/json")
	var book Books
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update A Book
func updateBook(w http.ResponseWriter, r * http.Request) {
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Books
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete A book
func deleteBook(w http.ResponseWriter, r * http.Request) {
w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index,item := range books {
		if item.ID == params["id"] {
			books = append(books[:index],books[index +1:]...)
			break

		}
		json.NewEncoder(w).Encode(books)
	}

}


func main()  {

	r := mux.NewRouter()

	books = append(books, Books{ID:"1",Isbn:"2400",Title:"Book One",Author:"Karim Mohamed"})
	books = append(books, Books{ID:"2",Isbn:"1000",Title:"Book Two",Author:"Ahmed Mohamed"})




	r.HandleFunc("/Books/api", getBooks).Methods("GET")
	r.HandleFunc("/Books/api/{id}", getBook).Methods("GET")
	r.HandleFunc("/Books/api", createBook).Methods("POST")
	r.HandleFunc("/Books/api/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/Books/api/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))
}
