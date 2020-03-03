# code-generator

### Local usage
```shell
go build generator.go
./generator
go clean
```

### Docker shortcut
`docker build -t trendev/code-generator . && docker push trendev/code-generator`

#### Run
`docker run -e CODE_SIZE=8 -it --rm trendev/code-generator`