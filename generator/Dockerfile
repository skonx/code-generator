# use goland:lastest instead of golang:alpine because go git is not available in alpine version
FROM golang as builder
WORKDIR /go/src/generator
COPY . .
RUN go get -v -d
RUN go build .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/generator/generator /app/
CMD ["./generator"]
