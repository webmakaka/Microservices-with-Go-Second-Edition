# Chapter 16 - Advanced Topics

* Static analysis of Go service code
* Implementing data validation
* Implementing streaming APIs
* Frameworks
* Storing microservice ownership data

<br/>

### Static analysis of Go service code

```
$ cd src
$ go vet ./...
```

**staticcheck**

```
$ go install honnef.co/go/tools/cmd/staticcheck@latest
$ staticcheck ./...
```

```
$ staticcheck -explain SA1019
```


```
https://staticcheck.dev/docs/checks/
https://staticcheck.dev/docs/running-staticcheck/ci/github-actions/
```

<br/>

### Implementing data validation


https://github.com/go-playground/validator


<br/>

### Implement streaming APIs

```
$ protoc -I=api --go_out=. --go-grpc_out=. movie.proto
```
