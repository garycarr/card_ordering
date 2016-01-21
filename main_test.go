package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testFixture struct {
	app *app

	// cleanup after the tests
	cleanup func()

	cardOrderingURL string
}

const defaultTestingMaxDeckSize = 5
const defaultTestingMaxShuffles = 50
const defaultTestingPrintEvery = 100
const defaultTestingStartDeckSize = 1

func createTestFixture(c Config) testFixture {
	fixture := testFixture{}
	c.isDev = true
	if c.maxDeckSize == 0 {
		c.maxDeckSize = defaultTestingMaxDeckSize
	}
	if c.maxShuffles == 0 {
		c.maxShuffles = defaultMaxShuffles
	}
	if c.printEvery == 0 {
		c.printEvery = defaultTestingPrintEvery
	}
	if c.startDeckSize == 0 {
		c.startDeckSize = defaultTestingStartDeckSize
	}

	c.port = 3001

	fixture.cleanup = func() {
		fixture.app.conf.isDev = false
		fixture.app.server.Close()
	}

	fixture.app = newApp(c)
	fixture.app.start()
	fixture.cardOrderingURL = fmt.Sprintf("http://%s", fixture.app.server.Addr)
	return fixture
}

// TestCardOrderingOverall runs the app until expDeckSize is reached.  It is possible but unlikely
// this test could take a long time.
// If you suspect the test is not working correctly then pass verbose:true into the config
func TestCardOrderingOverall(t *testing.T) {
	expDeckSize := 5
	tF := createTestFixture(Config{maxDeckSize: expDeckSize})
	tF.cleanup()
	assert.Equal(t, expDeckSize, tF.app.upToCard)
}

// Make sure that the defaults are applied
func TestGetConfigNoArgs(t *testing.T) {
	conf := getConfig()
	assert.Equal(t, conf.maxDeckSize, defaultMaxDeckSize)
	assert.Equal(t, conf.maxShuffles, defaultMaxShuffles)
	assert.Equal(t, conf.startDeckSize, defaultStartDeckSize)
	assert.Equal(t, conf.printEvery, defaultPrintEvery)
	assert.Equal(t, conf.verbose, defaultVerbose)
}

// Make sure that the defaults are applied
func TestGetConfigEnvVars(t *testing.T) {
	defer clearEnvs()
	os.Setenv("ENV_SET", "TRUE")
	os.Setenv("MAX_DECK_SIZE", "5")
	os.Setenv("MAX_SHUFFLES", "100")
	os.Setenv("PRINT_EVERY", "1000")
	os.Setenv("START_DECK_SIZE", "5")
	os.Setenv("VERBOSE", "false")

	conf := getConfig()
	assert.Equal(t, conf.maxDeckSize, 5)
	assert.Equal(t, conf.maxShuffles, 100)
	assert.Equal(t, conf.printEvery, 1000)
	assert.Equal(t, conf.startDeckSize, 5)
	assert.Equal(t, conf.verbose, false)
}

func clearEnvs() {
	os.Setenv("ENV_SET", "")
	os.Setenv("MAX_DECK_SIZE", "")
	os.Setenv("MAX_SHUFFLES", "")
	os.Setenv("PRINT_EVERY", "")
	os.Setenv("START_DECK_SIZE", "")
	os.Setenv("VERBOSE", "")
}
