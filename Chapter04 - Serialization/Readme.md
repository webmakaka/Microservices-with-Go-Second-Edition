# Chapter 04 - Serialization

Serialization is the process of converting data into a format that allows you to transfer it, store
it, and later deconstruct it.

Serialization has two primary use cases:

• Transferring data between services, acting as a common format between them  
• Encoding and decoding arbitrary data for storage, allowing you to store complex data
structures as byte arrays or regular strings


<br/>


// grpc
https://grpc.io/docs/protoc-installation/


<br/>

```
$ protoc -I=api --go_out=. movie.proto
```


<br/>


```
$ go mod tidy
```

<br/>

```
$ cd cmd/sizecompare
```

<br/>


```
$ go run main.go 
JSON size:	106B
XML size:	148B
Proto size:	63B
```


<br/>

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: movieexample.com/cmd/sizecompare
cpu: Intel(R) Core(TM) i5-4570 CPU @ 3.20GHz
BenchmarkSerializeToJSON-4    	 3477324	       463.0 ns/op
BenchmarkSerializeToXML-4     	  441832	      2516 ns/op
BenchmarkSerializeToProto-4   	 4292536	       315.5 ns/op
PASS
ok  	movieexample.com/cmd/sizecompare	4.743s
```
