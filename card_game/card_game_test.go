package main

import (
	"card_game/model"
	"card_game/stack"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	highDeal1 = []model.Card{
		model.Card{
			CardName:  "A",
			CardSuit:  "hearts",
			CardValue: 14,
		},
		model.Card{
			CardName:  "A",
			CardSuit:  "diamonds",
			CardValue: 14,
		},
		model.Card{
			CardName:  "A",
			CardSuit:  "clubs",
			CardValue: 14,
		},
	}
	highDeal2 = []model.Card{
		model.Card{
			CardName:  "K",
			CardSuit:  "hearts",
			CardValue: 13,
		},
		model.Card{
			CardName:  "k",
			CardSuit:  "diamonds",
			CardValue: 13,
		},
		model.Card{
			CardName:  "K",
			CardSuit:  "clubs",
			CardValue: 13,
		},
	}

	seqDeal1 = []model.Card{
		model.Card{
			CardName:  "J",
			CardSuit:  "hearts",
			CardValue: 11,
		},
		model.Card{
			CardName:  "Q",
			CardSuit:  "diamonds",
			CardValue: 12,
		},
		model.Card{
			CardName:  "K",
			CardSuit:  "clubs",
			CardValue: 13,
		},
	}

	seqDeal2 = []model.Card{
		model.Card{
			CardName:  "2",
			CardSuit:  "hearts",
			CardValue: 2,
		},
		model.Card{
			CardName:  "3",
			CardSuit:  "diamonds",
			CardValue: 3,
		},
		model.Card{
			CardName:  "4",
			CardSuit:  "clubs",
			CardValue: 4,
		},
	}

	pairDeal1 = []model.Card{
		model.Card{
			CardName:  "J",
			CardSuit:  "hearts",
			CardValue: 11,
		},
		model.Card{
			CardName:  "J",
			CardSuit:  "diamonds",
			CardValue: 11,
		},
		model.Card{
			CardName:  "K",
			CardSuit:  "clubs",
			CardValue: 13,
		},
	}

	pairDeal2 = []model.Card{
		model.Card{
			CardName:  "10",
			CardSuit:  "hearts",
			CardValue: 10,
		},
		model.Card{
			CardName:  "10",
			CardSuit:  "diamonds",
			CardValue: 10,
		},
		model.Card{
			CardName:  "9",
			CardSuit:  "clubs",
			CardValue: 9,
		},
	}

	topDeal1 = []model.Card{
		model.Card{
			CardName:  "A",
			CardSuit:  "hearts",
			CardValue: 14,
		},
		model.Card{
			CardName:  "10",
			CardSuit:  "diamonds",
			CardValue: 10,
		},
		model.Card{
			CardName:  "9",
			CardSuit:  "clubs",
			CardValue: 9,
		},
	}

	topDeal2 = []model.Card{
		model.Card{
			CardName:  "10",
			CardSuit:  "hearts",
			CardValue: 10,
		},
		model.Card{
			CardName:  "A",
			CardSuit:  "diamonds",
			CardValue: 14,
		},
		model.Card{
			CardName:  "9",
			CardSuit:  "clubs",
			CardValue: 9,
		},
	}

	random1 = []model.Card{
		model.Card{
			CardName:  "5",
			CardSuit:  "hearts",
			CardValue: 5,
		},
		model.Card{
			CardName:  "3",
			CardSuit:  "diamonds",
			CardValue: 3,
		},
		model.Card{
			CardName:  "7",
			CardSuit:  "clubs",
			CardValue: 7,
		},
	}

	random2 = []model.Card{
		model.Card{
			CardName:  "6",
			CardSuit:  "hearts",
			CardValue: 6,
		},
		model.Card{
			CardName:  "4",
			CardSuit:  "diamonds",
			CardValue: 4,
		},
		model.Card{
			CardName:  "2",
			CardSuit:  "clubs",
			CardValue: 2,
		},
	}

	random3 = []model.Card{
		model.Card{
			CardName:  "7",
			CardSuit:  "hearts",
			CardValue: 7,
		},
		model.Card{
			CardName:  "3",
			CardSuit:  "diamonds",
			CardValue: 3,
		},
		model.Card{
			CardName:  "8",
			CardSuit:  "clubs",
			CardValue: 8,
		},
	}
)

func GetDeckAndShuffle() stack.Deck {
	file, _ := ioutil.ReadFile("deck_of_cards.json")

	var deck stack.Deck
	err := json.Unmarshal([]byte(file), &deck)
	if err != nil {
		fmt.Printf("Error: Unmarshalling data %s\n", err.Error())
		return deck
	}
	deck.Shuffle()
	return deck
}

func TestFindWinnerHighDeal(t *testing.T) {
	deck := GetDeckAndShuffle()
	dealedCards := map[int][]model.Card{
		1: highDeal1,
		2: random1,
		3: random2,
		4: random3,
	}
	winner, reason := FindWinner(&deck, dealedCards)
	assert.Equal(t, 1, winner)
	assert.Equal(t, "High Deal", reason)
}

func TestFindWinnerSeqDeal(t *testing.T) {
	deck := GetDeckAndShuffle()
	dealedCards := map[int][]model.Card{
		1: seqDeal1,
		2: random1,
		3: random2,
		4: random3,
	}
	winner, reason := FindWinner(&deck, dealedCards)
	assert.Equal(t, 1, winner)
	assert.Equal(t, "Sequence Deal", reason)
}

func TestFindWinnerPairDeal(t *testing.T) {
	deck := GetDeckAndShuffle()
	dealedCards := map[int][]model.Card{
		1: random1,
		2: pairDeal2,
		3: random2,
		4: random3,
	}
	winner, reason := FindWinner(&deck, dealedCards)
	assert.Equal(t, 2, winner)
	assert.Equal(t, "Pair Deal", reason)
}

func TestFindWinnerTopCardDeal(t *testing.T) {
	deck := GetDeckAndShuffle()
	dealedCards := map[int][]model.Card{
		1: topDeal1,
		2: random1,
		3: random2,
		4: random3,
	}
	winner, reason := FindWinner(&deck, dealedCards)
	assert.Equal(t, 1, winner)
	assert.Equal(t, "Top Card Deal", reason)
}
