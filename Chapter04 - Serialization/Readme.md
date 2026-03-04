# Chapter04 - Serialization

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
