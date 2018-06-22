# appsync-lambda-sample

https://qiita.com/rioc/items/30699395e0080387f87e

## Build and Compression
```bash
$ GOOS=linux GOARCH=amd64 go build -o main
$ zip handler.zip ./main
```
