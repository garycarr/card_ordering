package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbout(t *testing.T) {
	tF := createTestFixture(Config{})
	defer tF.cleanup()
	res, err := http.Get(fmt.Sprintf("%s%s", tF.cardOrderingURL, "/about"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, res.StatusCode)
	aboutBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	assert.Contains(t, string(aboutBody), "This is a ongoing learning project")
}
