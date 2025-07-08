# Chapter07 - Storing Service Data

<br/>

```
$ docker run --name movieexample_db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=movieexample -p 3306:3306 -d mysql:latest
```

<br/>

```
$ cd src
```

<br/>

```
// OK!
$ docker exec -i movieexample_db mysql movieexample -h localhost -P 3306 --protocol=tcp -uroot -ppassword < schema/schema.sql
```


<br/>

```
// OK!
$ docker exec -i movieexample_db mysql movieexample -h localhost -P 3306 --protocol=tcp -uroot -ppassword -e "SHOW tables"
```

<br/>

```

$ cd rating/cmd/
$ go run *.go
```


<br/>


```
$ grpcurl -plaintext -d '{"record_id":"1", "record_type": "movie","user_id": "alex", "rating_value": 5}' localhost:8082 RatingService/PutRating
```

<br/>


```
// OK!
$ grpcurl -plaintext -d '{"record_id":"1", "record_type": "movie"}' localhost:8082 RatingService/GetAggregatedRating
```

<br/>

**response:**


```
{
  "ratingValue": 5
}
```