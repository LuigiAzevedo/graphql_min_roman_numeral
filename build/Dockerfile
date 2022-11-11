FROM golang:1.19

WORKDIR /usr

ADD src/ .
RUN go build -v -o /usr/local/bin/app cmd/server.go

CMD ["app"]

EXPOSE 8080