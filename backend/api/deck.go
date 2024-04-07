package api

import (
	"gorm.io/gorm"
	"net/http"
)

type DeckOption struct {
	gorm.Model
	// Deck Option은 Display option 이나
	//
}
type Algorithm interface {
	// gorm.Model
}
type Deck struct {
	gorm.Model
	DeckOption DeckOption
	Algorithm  Algorithm
	Words      []Word
}

func InitDeck(router *http.ServeMux) {

	router.HandleFunc("GET /api/deck/", handleGetAllDecks)
	router.HandleFunc("POST /api/deck/", handleCreateNewDeck)
	router.HandleFunc("GET /api/deck/{id}", handleGetDeckById)
	router.HandleFunc("UPDATE /api/deck/{id}", handleUpdateDeck)
}

func handleGetAllDecks(w http.ResponseWriter, r *http.Request) {

}

func handleGetDeckById(w http.ResponseWriter, r *http.Request) {

}

func handleCreateNewDeck(w http.ResponseWriter, r *http.Request) {
	// Deck Option, Algorithm
}

func handleUpdateDeck(w http.ResponseWriter, r *http.Request) {
	// Deck Option, Algorithm
}
