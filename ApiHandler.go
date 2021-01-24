package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)





func findDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	domain :=params["domain"]

	config, exist := dbContext.Find(domain)

	if exist {
		json.NewEncoder(w).Encode(config)
		return
	}
	json.NewEncoder(w).Encode(domain+" is not found!")
}

func InitApi(port int)  {

	router = mux.NewRouter()
	router.HandleFunc("/find/{domain}", findDomain).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
	fmt.Println("Listening port: "+strconv.Itoa(port))
}