FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 5003
CMD ["./main"]
