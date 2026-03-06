# Chapter 05 - Synchronous Communication


```
// grpc
https://grpc.io/docs/protoc-installation/
```

```
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

```
// grpcurl
https://github.com/fullstorydev/grpcurl

// Install
$ curl -sSL "https://github.com/fullstorydev/grpcurl/releases/download/v1.9.3/grpcurl_1.9.3_linux_x86_64.tar.gz" | sudo tar -xz -C /usr/local/bin
```


<br/>

```
// Made checkout . because of error
$ cd src
$ protoc -I=api --go_out=. --go-grpc_out=. movie.proto
```

<br/>

```
$ docker start dev-consul
```

<br/>

```
RUN-> all microservices
```

<br/>

```
$ grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie"}' localhost:8082 RatingService/GetAggregatedRating
```

<br/>

**response**

```
// OK!
ERROR:
  Code: NotFound
  Message: ratings not found for a record
```


<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
