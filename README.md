## Reverse string API / CLI

### A simple service to reorder a sentence (words) from the end, to the beginning

## How to run

* Build the project
```shell script
go build *.go
```

### Two type of ruuning interfaces
* Run as CLI
```shell script
go run main.go -s="enter some sentence here"
```

* Run as a web server
```shell script
curl --location --request POST 'localhost:9595/reverse?s=hello%20there'
```