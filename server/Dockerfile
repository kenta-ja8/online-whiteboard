FROM golang:1.17 AS build
WORKDIR /go/src
COPY go ./go
COPY app ./app
COPY main.go .
COPY go.mod .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
