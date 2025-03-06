FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
# RUN goose -dir migrations up
RUN go build -o main cmd/main.go
CMD ["./main"]