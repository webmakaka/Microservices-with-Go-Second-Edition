# Chapter05 - Synchronous Communication


```
// grpc
https://grpc.io/docs/protoc-installation/
```

```
// grpcurl
https://github.com/fullstorydev/grpcurl

// Install
$ curl -sSL "https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_x86_64.tar.gz" | sudo tar -xz -C /usr/local/bin
```


<br/>

```
// Made checkout . because of error
$ cd /src
$ protoc -I=api --go_out=. --go-grpc_out=. movie.proto
```


<br/>

```
RUN-> all microservices
$ docker start dev-consul
```

<br/>

```
$ grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie"}' localhost:8082 RatingService/GetAggregatedRating
```

<br/>

**response**

```
ERROR:
  Code: NotFound
  Message: ratings not found for a record
```