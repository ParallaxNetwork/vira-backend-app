FROM golang:1.21.4

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]