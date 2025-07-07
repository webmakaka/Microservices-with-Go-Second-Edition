# Chapter03

<br/>

```
// consul
$ docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client='0.0.0.0'
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
// OK!
http://localhost:8500/
```



```

$ curl -v 'localhost:8083/movie?id=1'
```


```
// movie service
2025/07/04 11:57:12 Calling metadata service. Request: GET http://localhost:8081/metadata
```