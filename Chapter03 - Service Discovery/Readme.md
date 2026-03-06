# Chapter 03 - Service Discovery

<br/>

```
// consul
$ docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client='0.0.0.0'
```



<br/>

```
$ cd src/metadata/cmd/
$ go run *.go

+1 instance
$ go run *.go --port 8084
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

<br/>

```
$ curl -v 'localhost:8083/movie?id=1'
```


<br/>

```
// movie service
2026/03/03 19:35:10 Calling metadata service. Request: GET http://localhost:8084/metadata
2026/03/03 19:35:18 Calling metadata service. Request: GET http://localhost:8081/metadata
2026/03/03 19:35:22 Calling metadata service. Request: GET http://localhost:8081/metadata
2026/03/03 19:35:24 Calling metadata service. Request: GET http://localhost:8081/metadata
2026/03/03 19:35:25 Calling metadata service. Request: GET http://localhost:8084/metadata
2026/03/03 19:35:25 Calling metadata service. Request: GET http://localhost:8084/metadata
2026/03/03 19:35:26 Calling metadata service. Request: GET http://localhost:8081/metadata
2026/03/03 19:35:27 Calling metadata service. Request: GET http://localhost:8084/metadata
2026/03/03 19:35:27 Calling metadata service. Request: GET http://localhost:8084/metadata
2026/03/03 19:35:28 Calling metadata service. Request: GET http://localhost:8081/metadata
2026/03/03 19:35:28 Calling metadata service. Request: GET http://localhost:8084/metadata
```


<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
