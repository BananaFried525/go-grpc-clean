
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go

EXPOSE 50051

CMD ["./server"]
