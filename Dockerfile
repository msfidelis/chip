FROM golang:1.15 AS builder

WORKDIR $GOPATH/src/chip

COPY . ./

RUN go get -u github.com/swaggo/swag/cmd/swag@v1.6.7

RUN go get -u

RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest

COPY --from=builder /go/src/chip/main ./

EXPOSE 8080

ENTRYPOINT ["./main"]
