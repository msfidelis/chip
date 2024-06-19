FROM fidelissauro/apko-go:latest-amd64 AS builder

WORKDIR /root/src/chip

COPY . ./

RUN go install github.com/swaggo/swag/cmd/swag@v1.7.8

RUN go get -u

RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM fidelissauro/apko-run:latest-amd64

COPY --from=builder /root/src/chip/main ./

EXPOSE 8080

ENTRYPOINT ["./main"]
