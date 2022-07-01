package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/willie/BookstoreAPI/src/routes"
)

func main() {

	//creating a new route instance
	r := mux.NewRouter()
	routes.RegirsterStoreRoutes(r)

	//handle the route
	http.Handle("/", r)

	//starting the server at port 3000
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
