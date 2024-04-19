FROM golang:1.22

RUN go install github.com/vektra/mockery/v2@latest 

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . .

WORKDIR /go/src/app/backend
RUN go build -o cook_droogers cmd/cookdroogers/main.go

ENV PORT=8080

# CMD ["./cook_droogers"]
CMD ["tail", "-f", "/dev/null"]
