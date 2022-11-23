FROM golang:1.19.3-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./ ./

RUN go mod tidy

CMD ["air"]
