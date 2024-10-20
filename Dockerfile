FROM golang:1.23 AS build

RUN go install github.com/vektra/mockery/v2@latest

WORKDIR /go/src/app

COPY . .

WORKDIR /go/src/app/backend
RUN go build -o cook_droogers cmd/cookdroogers/main.go
RUN go build -o rest-api internal/server/cmd/swagger-cook-droogers-server/main.go

WORKDIR  /app

RUN cp /go/src/app/backend/cook_droogers /app/cook_droogers && \
    cp /go/src/app/backend/config/config.yaml /app/config.yaml && \
    cp /go/src/app/backend/cmd/techUI/label_info.txt /app/label_info.txt && \
    cp /go/src/app/backend/Makefile /app/Makefile && \
    cp /go/src/app/backend/rest-api /app/rest-api

# #=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====

# FROM

# COPY --from=build /go/src/app/backend/cook_droogers /app/cook_droogers
# COPY --from=build /go/src/app/backend/config/config.yaml /app/config.yaml
# COPY --from=build /go/src/app/backend/cmd/techUI/label_info.txt /app/label_info.txt

ENV PORT=8080

WORKDIR /app

#CMD ["/app/rest-api", "--host", "0.0.0.0", "--port", "13337"]
#CMD ["tail", "-f", "/dev/null"]
