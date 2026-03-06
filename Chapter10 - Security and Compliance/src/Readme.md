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
$ cd metadata/configs/
$ go run ../cmd/*.go
```


<br/>

```
$ cd src/metadata/configs
```

<br/>

```
// OK!
$ grpcurl -cacert ca-cert.pem -cert metadata-cert.pem -key metadata-key.pem \
  -d '{
    "metadata": {
      "id": "alien-3",
      "title": "Alien 3",
      "description": "Ellen Ripley discovers that an unwanted guest came along for the ride.",
      "director": "David Fincher"
    }
  }' localhost:8081 MetadataService/PutMetadata
```

<br/>

```
// OK!
$ grpcurl -cacert ca-cert.pem \
  -cert metadata-cert.pem \
  -key metadata-key.pem \
  -d '{"movie_id":"alien-3"}' \
  localhost:8081 MetadataService/GetMetadata
```


**response:**

<br/>

```
{
  "metadata": {
    "id": "alien-3",
    "title": "Alien 3",
    "description": "Ellen Ripley discovers that an unwanted guest came along for the ride.",
    "director": "David Fincher"
  }
}
```

<br/>

**Need to update movie/cmd/main.go**

```
// FAIL!
$ grpcurl -cacert ca-cert.pem -d '{"movie_id":"1"}' localhost:8083 MovieService/GetMovieDetails
Failed to dial target host "localhost:8083": tls: first record does not look like a TLS handshake
```

```
// FAIL!
$ grpcurl -cacert ca-cert.pem \
  -cert metadata-cert.pem \
  -key metadata-key.pem \
  -d '{"movie_id":"alien-3"}' \
  localhost:8083 MovieService/GetMovieDetails
Failed to dial target host "localhost:8083": tls: first record does not look like a TLS handshake
```


<br/>

## Implementing authentication and access control with JWT


<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
