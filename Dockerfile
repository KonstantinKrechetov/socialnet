FROM golang:alpine

WORKDIR /socialnet
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd

CMD ["/socialnet/bin/api"]

EXPOSE 8080