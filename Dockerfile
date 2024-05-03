FROM golang:1.22 AS build

RUN go install github.com/vektra/mockery/v2@latest

WORKDIR /go/src/app

COPY . .

WORKDIR /go/src/app/backend
RUN go build -o cook_droogers cmd/cookdroogers/main.go

# #=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====

# FROM golang:1.22

# COPY --from=build /go/src/app/backend/cook_droogers /app/cook_droogers
# COPY --from=build /go/src/app/backend/config/config.yaml /app/config.yaml
# COPY --from=build /go/src/app/backend/cmd/techUI/label_info.txt /app/label_info.txt

ENV PORT=8080

# WORKDIR /app

# CMD ["./cook_droogers"]
CMD ["tail", "-f", "/dev/null"]
