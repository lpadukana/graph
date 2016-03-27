# Graph
Graphviz as a service

## Why
* Created to visualize the intermediate results of an algorithm
* HTTP and GET for easy use (in a browser)
* Golang because I am learning/liking it

## Local Use
* Install [graphviz](http://www.graphviz.org/)
* Run the program `go run main.go`
  * This will start the http server at port 9292
* Use the endpoint
  * `curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "digraph{a->{1,2,3}}"`
  * `curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "@elb.dot"`
