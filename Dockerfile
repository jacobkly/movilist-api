FROM golang:alpine

WORKDIR /klyvi-api
COPY go.mod go.sum ./

COPY . .

RUN go build -o ./bin/api ./cmd/api \
    && go build -o ./bin/migrate ./cmd/migrate

CMD ["/klyvi-api/bin/api"]
EXPOSE 8080