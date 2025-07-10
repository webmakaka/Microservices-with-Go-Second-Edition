# Chapter08 - Setting Up Service Deployments

<br/>

```
$ minikube start
$ eval $(minikube -p minikube docker-env)
```


<br/>


```
$ helm repo add hashicorp https://helm.releases.hashicorp.com

$ helm install consul hashicorp/consul --set global.name=consul
--create-namespace --namespace consul

$ kubectl port-forward service/consul-server 8500:8500 -n consul
```


<br/>


```
$ src/metadata
$ GOOS=linux go build -o main cmd/*.go
$ docker build -t metadata:latest .
$ kubectl apply -f kubernetes-deployment.yaml
```


<br/>


```
$ src/rating
$ GOOS=linux go build -o main cmd/*.go
$ docker build -t rating:latest .
$ kubectl apply -f kubernetes-deployment.yaml
```


<br/>


```
$ src/movie
$ GOOS=linux go build -o main cmd/*.go
$ docker build -t movie:latest .
$ kubectl apply -f kubernetes-deployment.yaml
```

<br/>

```
$ kubectl port-forward <POD_ID> 8081:8081
$ kubectl port-forward <POD_ID> 8082:8082
$ kubectl port-forward <POD_ID> 8083:8083
```


<br/>

```
$ grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie", "user_id": "alex", "rating_value": 5}' localhost:8082 RatingService/PutRating

$ grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie"}' localhost:8082 RatingService/GetAggregatedRating
```