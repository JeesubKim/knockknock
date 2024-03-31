package main

import (
	"fmt"
	"knockknock/middleware"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {

	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "hello world")
	})

	router.HandleFunc("GET /words", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "get words")
	})
	router.HandleFunc("GET /words/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		fmt.Fprintf(w, "get words of %s", id)
	})
	router.HandleFunc("POST /words", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "post words")
	})

	router.HandleFunc("GET /user/{id}", handleGetUserByID)

	server := http.Server{
		Addr:    s.addr,
		Handler: middleware.Logging(router),
	}
	fmt.Println("Server is up and running at ", s.addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Fprintf(w, "get user of %s", id)
}
