# Chapter09 - Unit and Integration Testing

```
$ go test
```

<br/>

## Mocking

```
$ go install github.com/golang/mock/mockgen@latest
```

```
$ cd src
$ mockgen -package=repository -source=metadata/internal/controller/metadata/controller.go
```

```
$ cd metadata/internal/controller/metadata/
$ go test
go: downloading github.com/stretchr/testify v1.7.1
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
PASS
ok  	movieexample.com/metadata/internal/controller/metadata	0.004s
```

<br/>

## Integration tests

Integration tests are automated tests that verify the correctness of integrations between the
individual units of your services and the services themselves.

```
$ go run test/integration/*.go
```
