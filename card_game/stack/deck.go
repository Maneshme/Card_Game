package stack

import (
	"card_game/model"
	"math/rand"
	"time"
)

type Deck []model.Card

// IsEmpty: check if stack is empty
func (deck *Deck) IsEmpty() bool {
	return len(*deck) == 0
}

// Push a new value onto the stack
func (deck *Deck) Push(str model.Card) {
	*deck = append(*deck, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (deck *Deck) Pop() (model.Card, bool) {
	if deck.IsEmpty() {
		return model.Card{}, false
	} else {
		index := len(*deck) - 1   // Get the index of the top most element.
		element := (*deck)[index] // Index into the slice and obtain the element.
		*deck = (*deck)[:index]   // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (deck *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(*deck); n > 0; n-- {
		randIndex := r.Intn(n)
		(*deck)[n-1], (*deck)[randIndex] = (*deck)[randIndex], (*deck)[n-1]
	}
}
