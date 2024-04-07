package api

import (
	"gorm.io/gorm"
	"knockknock/middleware"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Serve() {

	// 1. Create New Mux router
	router := http.NewServeMux()

	// 2. Mux all apis
	InitDeck(router)
	s.InitWord(router)

	// 3. Instantiate the server
	server := http.Server{
		Addr:    s.addr,
		Handler: middleware.Logging(router),
	}

	log.Println("Server is up and running at %s", s.addr)
	// 4. Run server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
