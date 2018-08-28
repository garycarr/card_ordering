package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// shuffleAndCheck continously loops, shuffling until the cards are in order
func (a *app) shuffleAndCheck(deckSize int) {
	//Create a deck of cards of n size
	deck := make([]int, deckSize)

	// Populate the deck
	i := 0
	for index := range deck {
		deck[index] = i
		i++
	}
	rand.Seed(time.Now().UnixNano())
	// Keep looping until we have a deck of cards shuffled in order
	for i := 0; i != a.conf.maxShuffles; i++ {
		// Shuffle the deck
		numShuffles := rand.Intn(10) + 10
		for count := 0; count < numShuffles; count++ {
			a.shuffle(deck)
		}
		if a.checkDeckOrder(deck) {
			if a.conf.verbose {
				log.Printf("Got %d cards in order after %d shuffles", deckSize, a.attempts)
			}
			// store how long it took to get to this point
			a.results[a.upToCard] = a.attempts
			a.upToCard++
			a.attempts = 0
			break
		}
		// Shuffle again as the deck was not in order
		if a.conf.verbose && a.attempts%a.conf.printEvery == 0 {
			log.Printf("Trying %d cards in order. So far %d shuffles", deckSize, a.attempts)
		}
		a.attempts++
	}
}

// Taken from http://marcelom.github.io/2013/06/07/goshuffle.html
func (a *app) shuffle(deck []int) {
	for range deck {
		i := rand.Intn(len(deck))
		j := rand.Intn(len(deck))
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func (a *app) checkDeckOrder(deck []int) bool {
	for i, card := range deck {
		if i != card {
			if a.conf.verbose {
				// Print out if we nearly had a matching deck
				if len(deck) > 5 && len(deck)-3 < i {
					fmt.Printf("Nearly matched on %d cards with %v on attempt %d\n", len(deck), deck, a.attempts)
				}
			}
			return false
		}
	}
	return true
}
