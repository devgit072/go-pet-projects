package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}

var books []Book

func init() {
	books = append(books,
		Book{ID:1, Title:"Golang", Author: "Hulk", Year: "2019"},
		Book{ID:2, Title:"Golang1", Author: "Hulk1", Year: "2020"},
		Book{ID:3, Title:"Golang2", Author: "Hulk2", Year: "2021"},
		Book{ID:4, Title:"Golang3", Author: "Hulk3", Year: "2022"},
	)
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting list of books")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Vars is map containing paramtereName and its value.
	// map[id]=3
	log.Println(vars)
	for _, book := range books {
		idVal,err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println("error: ", err)
		}

		if book.ID == idVal {
			json.NewEncoder(w).Encode(book)
			//json.NewEncoder(w).Encode(&book) even &book will also work.
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book) // read request and decode it and store it in variable book
	books = append(books, book)
	json.NewEncoder(w).Encode(&books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book) // read it in book.

	for i, b := range books {
		if b.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(&books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Error: ", err)
	}
	for i, book := range books {
		if id == book.ID {
			books = append(books[:i], books[i+1])
		}
	}

	json.NewEncoder(w).Encode(&books)
}