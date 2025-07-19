# Chapter13 - Setting Up Service Alerting


<br/>

```
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

<br/>

Something wrong for me!