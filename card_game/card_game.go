package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	model "card_game/model"
	stack "card_game/stack"
)

func main() {
	file, _ := ioutil.ReadFile("deck_of_cards.json")

	var deck stack.Deck
	players := 4
	err := json.Unmarshal([]byte(file), &deck)
	if err != nil {
		fmt.Printf("Error: Unmarshalling data %s\n", err.Error())
		return
	}
	deck.Shuffle()
	dealedCards := DealCards(&deck, players)
	fmt.Printf("Dealed Cards: %+v\n", dealedCards)
	winner, reason := FindWinner(&deck, dealedCards)
	fmt.Printf("Winner of this round is: %d due to %s\n", winner, reason)
}

// Function to find winner of round
func FindWinner(deck *stack.Deck, dealedCards map[int][]model.Card) (int, string) {
	winners := []int{}
	winners = FindHighDeal(dealedCards)
	if len(winners) != 0 {
		PrintWinningHands(dealedCards, winners)
		return AnalyzeWinners(deck, winners), "High Deal"
	}

	winners = FindSequences(dealedCards)
	if len(winners) != 0 {
		PrintWinningHands(dealedCards, winners)
		return AnalyzeWinners(deck, winners), "Sequence Deal"
	}

	winners = FindPair(dealedCards)
	if len(winners) != 0 {
		PrintWinningHands(dealedCards, winners)
		return AnalyzeWinners(deck, winners), "Pair Deal"
	}

	winners = FindTopCardHolder(dealedCards)
	if len(winners) != 0 {
		PrintWinningHands(dealedCards, winners)
		return AnalyzeWinners(deck, winners), "Top Card Deal"
	}

	return 0, "No Winner"
}

// Function to print winning hands
func PrintWinningHands(dealedCards map[int][]model.Card, winners []int) {
	for _, winner := range winners {
		fmt.Printf("Player: %d\n", winner)
		fmt.Printf("Winning Hand: %+v\n", dealedCards[winner])
	}
}

// Function to analyze winner of game
func AnalyzeWinners(deck *stack.Deck, winners []int) int {
	if len(winners) == 1 {
		return winners[0]
	}
	if len(winners) > 1 {
		return BreakTie(deck, winners)
	}
	return 0
}

// Function to break tie
func BreakTie(deck *stack.Deck, winners []int) int {
	fmt.Printf("Tie Between Players: %+v\n", winners)
	dealedCards := make(map[int][]model.Card)
	tieWinner := 0
	for _, winner := range winners {
		dealed := []model.Card{}
		card, isEmpty := deck.Pop()
		if !isEmpty {
			return 0
		}
		dealed = append(dealed, card)
		dealedCards[winner] = dealed
	}
	fmt.Printf("Dealed Cards To Break Tie %v\n", dealedCards)
	topPlayers := FindTopCardHolder(dealedCards)
	if len(topPlayers) > 1 {
		tieWinner = BreakTie(deck, topPlayers)
	} else {
		tieWinner = topPlayers[0]
	}
	fmt.Printf("Winning Card: %+v\n", dealedCards[tieWinner])
	return tieWinner
}

// Function to deal cards
func DealCards(deck *stack.Deck, players int) map[int][]model.Card {
	fmt.Printf("Players: %d\n", players)
	dealedCards := make(map[int][]model.Card)
	cardsToDeal := 3 * players
	player := 1
	for index := 0; index < cardsToDeal; index++ {
		dealed := []model.Card{}
		if val, ok := dealedCards[player]; ok {
			dealed = val
		}
		card, _ := deck.Pop()
		dealed = append(dealed, card)
		dealedCards[player] = dealed
		player++
		if player > players {
			player = 1
		}
	}
	return dealedCards
}

// Function to find if dealed cards are in sequence
func FindSequences(dealedCards map[int][]model.Card) []int {
	winners := []int{}
	for player, cards := range dealedCards {
		sort.SliceStable(cards, func(i, j int) bool {
			return cards[i].CardValue < cards[j].CardValue
		})
		if cards[0].CardValue+1 == cards[1].CardValue &&
			cards[1].CardValue+1 == cards[2].CardValue {
			winners = append(winners, player)
		} else if (cards[2].CardValue == 14 && cards[0].CardValue == 2) &&
			(cards[1].CardValue == 3 || cards[1].CardValue == 13) {
			winners = append(winners, player)
		}
	}
	return winners
}

// Function to find if dealed cards are a high deal
func FindHighDeal(dealedCards map[int][]model.Card) []int {
	winners := []int{}
	for player, cards := range dealedCards {
		if cards[0].CardValue == cards[1].CardValue &&
			cards[1].CardValue == cards[2].CardValue {
			winners = append(winners, player)
		}
	}
	return winners
}

func FindPair(dealedCards map[int][]model.Card) []int {
	winners := []int{}
	for player, cards := range dealedCards {
		if cards[0].CardValue == cards[1].CardValue ||
			cards[1].CardValue == cards[2].CardValue ||
			cards[0].CardValue == cards[2].CardValue {
			winners = append(winners, player)
		}
	}
	return winners
}

func FindTopCardHolder(dealedCards map[int][]model.Card) []int {
	topCard := 0
	topPlayer := []int{}
	for player, cards := range dealedCards {
		for _, card := range cards {
			if card.CardValue > topCard {
				topCard = card.CardValue
				topPlayer = []int{player}
			} else if card.CardValue == topCard {
				if !intInSlice(player, topPlayer) {
					topPlayer = append(topPlayer, player)
				}
			}
		}
	}
	return topPlayer
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
