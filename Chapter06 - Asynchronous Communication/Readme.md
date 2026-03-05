# Chapter06 - Asynchronous Communication

Asynchronous communication is communication between a sender and one or multiple receivers,
where a sender does not necessarily expect an immediate response to their messages.

<br/>

```
$ docker start dev-consul
```

<br/>

```
$ cd src/cmd/ratingproducer/
$ docker-compose up -d
```

<br/>

```
$ docker exec -it kafka /bin/bash
```

<br/>

```
# cd /opt/kafka/bin
# kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic ratings
```

<br/>

+ New Terminal

```
$ cd /src/rating/cmd
$ go run main.go
```

<br/>

+ New Terminal

```
$ cd /src/cmd/ratingproducer
$ go run main.go
```

<br/>

```
// Rating service
// OK!
Consumed a message: {{1 movie 105 5} test-provider put}
Consumed a message: {{2 movie 105 4} test-provider put}
```
