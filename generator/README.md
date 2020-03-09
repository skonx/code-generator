# code-generator

### Local usage
```shell
go build -o generator main.go
./generator
go clean
```

### Docker shortcut
`docker build -t trendev/code-generator . && docker push trendev/code-generator`

#### Run
`docker build -t trendev/code-generator . && docker run -it --rm trendev/code-generator`

`docker run -e CODE_SIZE=8 -e DELAY=100 -it --rm trendev/code-generator`