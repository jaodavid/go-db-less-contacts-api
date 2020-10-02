package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Contact struct (Model)
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Init contacts var as a slice Contact struct
var contacts []Contact

// Get all contacts
func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// Get single contact
func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Looping through contacts and find one with the id from the params
	for _, item := range contacts {
		if item.ID == convertStrIDtoInt(params["id"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

// Add new contact
func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = getMaxID() + 1
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// Update contact
func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.ID == convertStrIDtoInt(params["id"]) {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.ID = convertStrIDtoInt(params["id"])
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
}

// Delete contact
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.ID == convertStrIDtoInt(params["id"]) {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Get Max ID from Contacts
func getMaxID() int {
	max := 0
	for _, item := range contacts {
		if item.ID > max {
			max = item.ID
		}
	}
	return max
}

func convertStrIDtoInt(id string) int {
	i, _ := strconv.ParseInt(id, 10, 32)
	return int(i)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	contacts = append(contacts, Contact{ID: 0, Name: "Jao", Phone: "09956100541", Email: "markjourine.david@gmail.com"})
	contacts = append(contacts, Contact{ID: 1, Name: "Shanie", Phone: "96xxx-xxxxx", Email: "shanie.baladjay@gmail.com"})
	contacts = append(contacts, Contact{ID: 2, Name: "John", Phone: "97xxx-xxxxx", Email: "jhon.doe@gmail.com"})

	// Route handles & endpoints
	r.HandleFunc("/contacts", getContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts/{id}", updateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}", deleteContact).Methods("DELETE")

	// Start server
	fmt.Println("Port 3000 is running...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
