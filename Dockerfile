FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main/main.go

EXPOSE 8081

CMD [ "./main" ]
