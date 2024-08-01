FROM golang:1.22-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY cmd .
COPY web .

RUN go build -v -o urs/local/bin/app ./cmd

CMD ["app"]
