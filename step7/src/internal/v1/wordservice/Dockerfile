FROM golang:1.13.12-stretch as builder

WORKDIR /go/src/application-modernization-example
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod download
RUN go build -o app main.go

FROM scratch
COPY --from=builder /go/src/application-modernization-example/app /app

EXPOSE 50000
ENTRYPOINT ["/app"]
