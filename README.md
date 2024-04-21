### a simple round robin balancer

### installation

- install [go](go.dev)
- run ```go mod tidy``` for installing dependencies

### start a few servers

- ```go run server.go -port 8081```
- ```go run server.go -port 8082```

### start the load balancer

- cd balancer
- ```go run balancer.go```

### send requests to the balancer

- ```curl 127.0.0.1:8000/hello```
- requests are sent to the servers in a linear way
