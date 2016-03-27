# Graph
Graphviz as a service

## Why
* Created to visualize the intermediate results of an algorithm
* HTTP and GET for easy use (embed in webpages with CORS)
* Golang because I am learning/liking it

## For development
* Install [graphviz](http://www.graphviz.org/)
* Run the program `go run main.go`
  * This will start the http server at port 9292
* Use the endpoint
  * `curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "digraph{a->{1,2,3}}"`
  * `curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "@examples/service.dot"`
  * `curl -G -v "http://localhost:9292/dot/to/png" --data-urlencode "@examples/by-two.dot"`

## Build
```sh
$ GOOS=linux go build
$ GRAPH_VERSION=1
$ docker build -t graph:${GRAPH_VERSION} .
```

## Run
```sh
$ docker run --rm -it -p 9292:9292 graph:${GRAPH_VERSION}
```
