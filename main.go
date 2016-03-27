package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/dot/to/svg", func(rw http.ResponseWriter, r *http.Request) {
		query, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			http.Error(rw, "Bad Request", 400)
		}

		dotCommand := exec.Command("dot", "-Tsvg")
		dotCommand.Stdin = strings.NewReader(query)
		dotCommand.Stdout = rw
		dotCommand.Stderr = os.Stderr

		err = dotCommand.Run()
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}
	})

	log.Println(http.ListenAndServe(":9292", nil))
}
