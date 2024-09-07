# Step 1: Setup
FROM golang:latest
WORKDIR /app

COPY go.* ./
COPY cmd cmd
COPY internal internal

RUN go mod download && go mod verify

RUN go build -o /application cmd/songlinkr/main.go

CMD ["/application"]
