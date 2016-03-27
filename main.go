package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func renderDot(rw http.ResponseWriter, r *http.Request, format string) {
	query, err := url.QueryUnescape(r.URL.RawQuery)
	if err != nil {
		http.Error(rw, "Bad Request", 400)
	}

	dotCommand := exec.Command("dot", "-T"+format)
	dotCommand.Stdin = strings.NewReader(query)
	dotCommand.Stdout = rw
	dotCommand.Stderr = os.Stderr

	err = dotCommand.Run()
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
}

func main() {
	toFormats := []string{"svg", "png", "plain"}
	for _, toFormat := range toFormats {
		// need a local copy or the format will point to the last format
		var format = toFormat
		http.HandleFunc("/dot/to/"+format, func(rw http.ResponseWriter, r *http.Request) {
			renderDot(rw, r, format)
		})
	}
	log.Println(http.ListenAndServe(":9292", nil))
}
