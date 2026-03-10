# Chapter 13 - Setting Up Service Alerting

<br/>

https://sre.google/sre-book/practical-alerting/

<br/>

```
// http://localhost:8500/
$ docker start dev-consul

$ docker start jaeger
```

<br/>

```
$ cd src
$ docker run \
  -p 9090:9090 \
  -v ./configs:/etc/prometheus \
  prom/prometheus
```

<br/>


<br/>

Run all micros

```
// OK!
http://localhost:9091/metrics

// OK!
http://localhost:9093/metrics
```

<br/>

```
// OK!
http://localhost:9090/
```

<br/>

```
>_ up 

Execute
```

<br/>


```
$ cd src
$ docker run -p 9093:9093 -v ./configs:/etc/alertmanager prom/alertmanager --config.file=/etc/alertmanager/alertmanager.yml
```

```
http://localhost:9093/#/alerts
```


Something wrong for me!


<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
