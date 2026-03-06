# Chapter09 - Unit and Integration Testing

```
$ go test
```

<br/>

## Mock

```
$ go install github.com/golang/mock/mockgen@latest
```

```
$ cd src
$ mockgen -package=repository -source=metadata/internal/controller/metadata/controller.go
```
