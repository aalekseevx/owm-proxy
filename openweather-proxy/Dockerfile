FROM golang:1.20

WORKDIR /usr/src/openweather-proxy

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/openweather-proxy openweather-proxy/openweather-proxy.go

CMD ["openweather-proxy"]
