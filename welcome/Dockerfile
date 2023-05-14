FROM golang:1.20

WORKDIR /usr/src/welcome

COPY go.mod go.sum .
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/welcome welcome/welcome.go

CMD ["welcome"]
