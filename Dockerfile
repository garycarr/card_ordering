FROM golang:1.8

WORKDIR /go/src/github.com/garycarr/card_ordering
COPY . /go/src/github.com/garycarr/card_ordering

RUN go-wrapper install

CMD ["go-wrapper", "run"]
