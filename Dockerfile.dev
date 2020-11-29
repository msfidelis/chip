FROM golang:1.15 AS builder

# Install Air
RUN go get -u github.com/cosmtrek/air

WORKDIR $GOPATH/src/chip

COPY . ./

RUN pwd; ls -lha

CMD ["air"]