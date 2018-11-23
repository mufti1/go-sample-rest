package main

import (
	"log"
	"net/http"

	"github.com/mufti1/simple-rest/handlers"

	"github.com/gorilla/mux"
	"github.com/mufti1/simple-rest/db"
)

const (
	server = ":9000"
)

// Handler is abstraction of handler function
type Handler interface {
	GetPokedex(w http.ResponseWriter, r *http.Request)
	InsertPokedex(w http.ResponseWriter, r *http.Request)
}

// NewRouter implement function of router list
func NewRouter(handler Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/pokedex", handler.GetPokedex).Methods("GET")
	r.HandleFunc("/pokedex", handler.InsertPokedex).Methods("POST")

	return r
}

func main() {
	mongo, err := db.NewMongoDB("54.201.248.209:27017", "free", "root", "gampangdibobol", 1)
	if err != nil {
		log.Fatalf("errror mongodb connection: %v", err)
	}
	handler := handlers.NewHandler(mongo)
	r := NewRouter(handler)
	log.Printf("server running on %v", server)
	err = http.ListenAndServe(server, r)
	if err != nil {
		log.Fatalf("Unable to run http server: %v", err)
	}
	log.Println("Stopping API Service...")
}
