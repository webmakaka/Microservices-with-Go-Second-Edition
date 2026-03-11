# Chapter 14 - Performance Monitoring

```
$ docker run --add-host host.docker.internal:host-gateway -p 3000:3000 grafana/grafana-oss
```

```
// admin / admin
http://localhost:3000
```

Connections -> Data sources -> Add data source -> Prometheus

Prometheus server URL: http://host.docker.internal:9090

Save and test


Dashboards page -> Create dashboard -> Add visualization -> Prometheus data source -> Metric field -> process_open_fds -> Run queries button -> Save dashboard

same:  Metric: go_gc_duration_seconds


```
$ grpcurl -plaintext -d '{}' localhost:8081 MetadataService/GetMetadata
$ grpcurl -plaintext -d '{"movie_id": "0"}' localhost:8081 MetadataService/GetMetadata
```

http://localhost:9090

```
>_ call
```

Metric: call

<br/>

## Profiling Go services

<br/>

### Profiling CPU usage using the pprof tool

https://graphviz.org/


```
$ go run *.go --simulatecpuload
```


```
$ go tool pprof http://localhost:6060/debug/pprof/profile?seconds=5

$ go tool pprof -raw -output=cpu.prof 'http://localhost:6060/debug/pprof/profile?seconds=10'
```


```
$ perl stackcollapse-go.pl cpu.prof | perl flamegraph.pl > flame.svg
```

<br/>

### Profiling heap memory usage using the pprof tool

```
$ go tool pprof http://localhost:6060/debug/pprof/heap
```

<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
