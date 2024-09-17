FROM golang:1.23.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /quiz-app cmd/quiz/main.go

EXPOSE 8080

CMD ["/quiz-app"]
