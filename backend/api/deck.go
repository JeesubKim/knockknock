package main

import "http"

type DeckOption struct {
	// Deck Option은 Display option 이나
	//
}
type Algorithm interface {
}
type Deck struct {
	DeckOption DeckOption
	Algorithm  Algorithm
	Words      []Word
}

func InitDeck(router http.Handler) {

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
