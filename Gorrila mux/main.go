package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Index</h1>"))
	})

	routerAdmin := router.PathPrefix("/admin").Subrouter()
	{
		routerAdmin.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>index</h1>"))
		})
		routerAdmin.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>login</h1>"))
		})
		routerAdmin.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>register</h1>"))
		})
	}

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
