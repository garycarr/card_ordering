# card_ordering

Mess about and throw away Golang test for http://qi.com/infocloud/playing-cards. How many shuffles does it take to get cards in a correct numerical order.

# To run

## Command Line

To see the flags and descriptions run  `go install && card_ordering --help`

## Docker

In docker set the env vars in the compose file and run
```
docker-compose build && docker-compose up
```
or

```
docker build -t card_ordering .
docker run -e ENV_SET=true \
    -e MAX_DECK_SIZE=52 \
    -e MAX_SHUFFLES=-1 \
    -e PRINT_EVERY=100000 \
    -e START_DECK_SIZE=5 \
    -e VERBOSE=true \
    --name card_ordering
    card_ordering
```

After running, go to localhost:3000 to see the status in a browser
# sudoku
