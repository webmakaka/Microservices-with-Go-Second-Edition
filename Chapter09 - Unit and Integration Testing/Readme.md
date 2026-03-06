# Chapter 09 - Unit and Integration Testing

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

```
$ cd src
$ go run test/integration/*.go
go: downloading github.com/google/go-cmp v0.5.6
2026/03/06 14:44:12 Starting the integration test
2026/03/06 14:44:12 Setting up service handlers and clients
2026/03/06 14:44:12 Starting metadata service on localhost:8081
2026/03/06 14:44:12 Starting rating service on localhost:8082
2026/03/06 14:44:12 Starting movie service on localhost:8083
2026/03/06 14:44:12 Saving test metadata via metadata service
2026/03/06 14:44:12 Retrieving test metadata via metadata service
2026/03/06 14:44:12 Getting movie details via movie service
2026/03/06 14:44:12 Saving first rating via rating service
2026/03/06 14:44:12 Retrieving initial aggregated rating via rating service
2026/03/06 14:44:12 Saving second rating via rating service
2026/03/06 14:44:12 Saving new aggregated rating via rating service
2026/03/06 14:44:12 Getting updated movie details via movie service
2026/03/06 14:44:12 Integration test execution successful
```

<br/>

## code coverage

```
$ go test -cover ./metadata/internal/controller/metadata/
ok  	movieexample.com/metadata/internal/controller/metadata	0.005s	coverage: 83.3% of statements
```

```
$ go test -cover ./...
```


<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
