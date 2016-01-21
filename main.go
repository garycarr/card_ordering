package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

const defaultMaxDeckSize = 100
const defaultMaxShuffles = -1
const defaultPrintEvery = 1000000
const defaultStartDeckSize = 1
const defaultVerbose = true
const defaultPort = 3000

var (
	maxDeckSize   = flag.Int("max-deck-size", defaultMaxDeckSize, "Maximum number of cards in deck")
	maxShuffles   = flag.Int("max-shuffles", defaultMaxShuffles, "The max shuffles to try per deck size.  Input -1 to try forever")
	printEvery    = flag.Int("print-every", defaultPrintEvery, "How many shuffles to print at")
	startDeckSize = flag.Int("start-deck-size", defaultStartDeckSize, "The deck size to start with")
	verbose       = flag.Bool("verbose", defaultVerbose, "Prints out at every print-every")
)

func getConfig() Config {
	// Should be able to make this neater
	var conf Config
	if os.Getenv("ENV_SET") != "" {
		// Find a package to do this
		// Need to fail if not found
		var verbose bool
		if os.Getenv("VERBOSE") == "true" {
			verbose = true
		}
		conf = Config{
			maxDeckSize:   mustGetEnv("MAX_DECK_SIZE"),
			maxShuffles:   mustGetEnv("MAX_SHUFFLES"),
			startDeckSize: mustGetEnv("START_DECK_SIZE"),
			printEvery:    mustGetEnv("PRINT_EVERY"),
			verbose:       verbose,
		}
	} else {
		flag.Parse()
		conf = Config{
			maxDeckSize:   *maxDeckSize,
			maxShuffles:   *maxShuffles,
			printEvery:    *printEvery,
			startDeckSize: *startDeckSize,
			verbose:       *verbose,
		}
	}
	conf.port = defaultPort
	// TODO - sort this
	// 	conf.isDev = true
	return conf
}

func main() {
	app := newApp(getConfig())
	app.start()
}

func mustGetEnv(env string) int {
	envInt, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		log.Fatalf("Unable to get envVar for %s. It must be set", env)
	}
	return envInt
}
