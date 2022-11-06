package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type testUser struct {
	Name string
	Age  string
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user testUser
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)

		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		fmt.Printf("User name")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", mainLogic).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
