FROM golang:1.21

WORKDIR /usr/src/app

COPY App/go.mod App/go.sum ./
RUN go mod download && go mod verify

COPY App/ .

EXPOSE 8000 8000

RUN rm -f /usr/src/app/app
RUN go build -v -o /usr/src/app ./...
CMD ["go", "run", "main.go"]