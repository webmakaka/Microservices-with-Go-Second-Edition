# Chapter05 - Synchronous Communication

<br/>

```
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

<br/>

```
// grpcurl tool
https://github.com/fullstorydev/grpcurl
```


<br/>


```
$ protoc -I=api --go_out=. --go-grpc_out=. movie.proto
```