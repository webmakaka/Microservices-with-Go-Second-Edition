# Chapter 10 - Security and Compliance

<br/>

```
$ docker start dev-consul
```

<br/>

```
$ sudo vi /etc/hosts
127.0.0.1 consul-server.consul.svc.cluster.local
```

<br/>

```
$ cd src/metadata/configs
```

<br/>

```
// It makes metadata-key.pem and metadata-cert.pem in the current location
$ openssl req -x509 -nodes -newkey rsa:4096 \
  -keyout metadata-key.pem -out metadata-cert.pem -days 365 -nodes \
  -subj "/C=US/ST=State/L=City/O=Organization/OU=Department/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,DNS:example.com,IP:127.0.0.1,IP:192.168.1.1" -config /dev/null
```

```
$ cp metadata-cert.pem ca-cert.pem
```

Move the files that were generated in the previous step into the configs directory of each
of our microservices.


```
$ src/rating/configs
$ go run ../cmd/*.go
```

```
$ cd metadata/configs/
$ go run ../cmd/*.go
```

```
$ cd src/movie/configs/
$ go run ../cmd/*.go
```


<br/>

```
$ cd src/metadata/configs
```

```
// FAIL!
$ grpcurl -cacert ca-cert.pem -d '{"movie_id":"1"}' localhost:8081 MetadataService/GetMetadata
Failed to dial target host "localhost:8081": context deadline exceeded
```



```
$ grpcurl -cacert ca-cert.pem -d '{"movie_id":"1"}' localhost:8083 MovieService/GetMovieDetails
```

<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
