package model

// Represents an individual card in a deck
type Card struct {
	CardName  string `json:"cardName"`
	CardSuit  string `json:"cardSuit"`
	CardValue int    `json:"cardValue"`
}
