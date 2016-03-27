# Graph
Graphviz as a service

## Local Use
* Install graphviz
* Run the program `go run main.go`
* Examples
  * `curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "digraph{a->{1,2,3}}"`
  * `curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "@elb.dot"`
