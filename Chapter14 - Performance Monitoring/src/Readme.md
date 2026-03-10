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

same: go_gc_duration_seconds


<br/><br/>

---

<br/>

**Marley**  
<a href="https://k8s.ru/">Предложить инженеру работу / подработку на проекте с kubernetes, microservices, golang</a>
