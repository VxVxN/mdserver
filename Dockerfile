FROM arm32v7/golang:latest

ENV GOFLAGS="-mod=vendor" GO111MODULE=on
ENV GOARCH=arm

ADD . /app
WORKDIR /app

RUN go build -o app/mdserver ./cmd

RUN cd app

CMD ["./app/mdserver"]
