# Chapter02 - Scaffolding a Go Microservice


<br/>

```
• Metadata service: localhost:8081
• Rating service: localhost:8082
• Movie service: localhost:8083
```

<br/>

```
$ cd src/metadata/cmd/
$ go run *.go
```

<br/>


```
$ cd src/movie/cmd/
$ go run *.go
```

<br/>


```
$ cd src/rating/cmd/
$ go run *.go
```

<br/>

```
// metadata service
// 404
$ curl 'localhost:8081?id=1'
```


<br/>


```
// API service
// 404
$ curl 'localhost:8082?id=1&type=2'
```


<br/>


```
// movie service
// 404
$ curl 'localhost:8081?id=1'
```