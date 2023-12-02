FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8000 8000

RUN rm -f /usr/src/app/app
RUN go build -v -o /usr/src/app ./...
CMD ["go", "run", "main.go"]