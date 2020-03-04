FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go build -o generator .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/generator /app/
CMD ["./generator"]
