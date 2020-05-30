FROM golang:1.13 AS builder

# Install Air
RUN curl -fLo air https://git.io/linux_air
RUN mv air /usr/bin/air && chmod +x /usr/bin/air

WORKDIR $GOPATH/src/chip

COPY . ./

RUN go get -u

RUN pwd; ls -lha

CMD ["/usr/bin/air"]