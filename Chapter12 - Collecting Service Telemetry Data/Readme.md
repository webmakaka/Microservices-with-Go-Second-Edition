# Chapter12 - Collecting Service Telemetry Data


```
$ docker run -d --name jaeger \
-e COLLECTOR_OTLP_ENABLED=true \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 4317:4317 \
-p 4318:4318 \
-p 14250:14250 \
-p 14268:14268 \
-p 14269:14269 \
-p 9411:9411 \
jaegertracing/all-in-one:1.37
```


<br/>

```
$ cd src/metadata/cmd/
$ cp ../configs/default.yaml .
$ vi default.yaml
```

<br/>


```
 consul:
    address: http://localhost:8500
```

<br/>

```
$ go run *.go
```




<br/>


```
$ cd src/movie/cmd/
$ cp ../configs/default.yaml .
$ go run *.go
```

<br/>


```
$ cd src/rating/cmd/
$ cp ../configs/default.yaml .
$ go run *.go
```


<br/>


```
$ grpcurl -plaintext -d '{"movie_id":"1"}' localhost:8083 MovieService/GetMovieDetails
```

<br/>

```
// OK!
http://localhost:16686/
```


<br/>

```
service -> movie -> Find Traces
```