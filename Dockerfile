FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o main .
RUN go mod tidy
EXPOSE 8080
CMD ["./main"]