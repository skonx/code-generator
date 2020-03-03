FROM golang:alpine
WORKDIR /go/src/code-generator
COPY . .

RUN go get -d -v .
RUN go install -v .

CMD ["code-generator"]
