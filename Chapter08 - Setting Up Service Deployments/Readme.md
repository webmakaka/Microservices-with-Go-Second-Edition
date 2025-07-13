# Chapter08 - Setting Up Service Deployments

<br/>

```
$ minikube start
$ eval $(minikube -p minikube docker-env)
```


<br/>


```
// $ helm repo add hashicorp https://helm.releases.hashicorp.com

// FAIL!
// $ helm install consul hashicorp/consul --set global.name=consul --create-namespace --namespace consul
// $ helm pull hashicorp/consul --untar
```

<br/>

```
$ cd consul
$ helm install consul . --set global.name=consul --create-namespace --namespace consul
```

<br/>

```
$ kubectl get pods -n consul
NAME                                          READY   STATUS    RESTARTS   AGE
consul-connect-injector-78fdcf77d6-zzmnv      1/1     Running   0          45s
consul-server-0                               1/1     Running   0          44s
consul-webhook-cert-manager-d8c8f756b-kcmd5   1/1     Running   0          45s
```

<br/>


```
$ src/metadata
$ GOOS=linux CGO_ENABLED=0 go build -o main cmd/*.go
$ docker build -t metadata:latest .
$ kubectl apply -f kubernetes-deployment.yaml
```


<br/>


```
$ src/rating
$ GOOS=linux CGO_ENABLED=0 go build -o main cmd/*.go
$ docker build -t rating:latest .
$ kubectl apply -f kubernetes-deployment.yaml
```


<br/>


```
$ src/movie
$ GOOS=linux CGO_ENABLED=0 go build -o main cmd/*.go
$ docker build -t movie:latest .
$ kubectl apply -f kubernetes-deployment.yaml
```


<br/>

```
$ kubectl port-forward service/consul-server 8500:8500 -n consul
```


```
// OK!
http://localhost:8500/ui/dc1/services
```

<br/>

```
$ kubectl port-forward metadata 8081:8081
$ kubectl port-forward rating 8082:8082
$ kubectl port-forward movie 8083:8083
```


<br/>

```
// OK!
$ grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie", "user_id": "alex", "rating_value": 5}' localhost:8082 RatingService/PutRating

// OK!
$ grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie"}' localhost:8082 RatingService/GetAggregatedRating
```

<br/>

**response**

```
{
  "ratingValue": 5
}
```
