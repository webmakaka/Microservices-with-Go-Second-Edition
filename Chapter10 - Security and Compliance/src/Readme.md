# Chapter 10 - Security and Compliance

```
$ openssl req -x509 -nodes -newkey rsa:4096 \
  -keyout server.key -out server.crt -days 365 -nodes \
  -subj "/C=US/ST=State/L=City/O=Organization/OU=Department/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,DNS:example.com,IP:127.0.0.1,IP:192.168.1.1" -config /dev/null
```

Move the files that were generated in the previous step into the configs directory of each
of our microservices.

```
$ go run ../cmd/*.go
```


```
$ grpcurl -cacert server.crt -d '{"movie_id":"1"}' localhost:8081 MetadataService/GetMetadata
```

<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
