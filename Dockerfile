FROM golang:alpine

WORKDIR /movilist-api
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/movilist-api/bin/api"]
EXPOSE 8080