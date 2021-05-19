FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main ./server.go
EXPOSE 5003
CMD ["./main"]
