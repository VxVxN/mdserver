FROM golang:1.16

ENV GOFLAGS="-mod=vendor" GO111MODULE=on

ADD . /app
WORKDIR /app

RUN go build -o app/mdserver ./cmd

RUN cd app

CMD ["./app/mdserver"]
