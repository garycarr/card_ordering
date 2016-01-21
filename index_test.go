package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	tF := createTestFixture(Config{})
	defer tF.cleanup()
	res, err := http.Get(fmt.Sprintf("%s%s", tF.cardOrderingURL, "/"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, res.StatusCode)
	indexBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	assert.Contains(t, string(indexBody), "Card status")
}

func TestNotFound(t *testing.T) {
	tF := createTestFixture(Config{})
	defer tF.cleanup()
	res, err := http.Get(fmt.Sprintf("%s%s", tF.cardOrderingURL, "/foo"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
