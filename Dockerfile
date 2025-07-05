FROM golang:1.23-alpine

WORKDIR /app

COPY . ./

RUN go mod tidy && go mod download && go build -o trading-system

EXPOSE 8080
CMD ["./trading-system"]
