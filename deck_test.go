package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDeckOrder(t *testing.T) {
	type testData struct {
		description string
		deck        []int
		expBool     bool
	}
	testTable := []testData{
		testData{
			description: "deck 1 length, in order",
			deck:        []int{0},
			expBool:     true,
		},
		testData{
			description: "deck 3 length, in order",
			deck:        []int{0, 1, 2},
			expBool:     true,
		},
		testData{
			description: "deck 3 length, out of order",
			deck:        []int{0, 2, 1},
			expBool:     false,
		},
	}
	tF := createTestFixture(Config{})
	tF.cleanup()
	for _, td := range testTable {
		assert.Equal(t, td.expBool, tF.app.checkDeckOrder(td.deck), td.description)
	}
}
