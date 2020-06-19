# use goland:lastest instead of golang:alpine because go git is not available in alpine version
FROM golang as builder
WORKDIR /go/src/app
COPY . .
RUN go get -v -d
RUN go build -o generator .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/generator /app/
CMD ["./generator"]
