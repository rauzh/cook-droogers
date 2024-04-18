# Используем официальный образ Golang 1.22
FROM golang:1.22

RUN go get github.com/vektra/mockery/.../

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . .

RUN go build -o cook_droogers backend/cmd/cookdroogers/main.go

ENV PORT=8080

CMD ["./cook_droogers"]