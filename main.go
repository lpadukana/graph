package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/dot/to/svg", func(rw http.ResponseWriter, r *http.Request) {
		query, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			http.Error(rw, "Bad Request", 400)
		}

		io.WriteString(rw, query)
	})

	log.Println(http.ListenAndServe(":9292", nil))
}
