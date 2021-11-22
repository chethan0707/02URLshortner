package main

import (
	"URLshortner/urlshortner"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	mapHandler := urlshortner.MapHandler(pathToUrls, r)
	yamlHandler, err := urlshortner.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		fmt.Println(err)
	}
	r.HandleFunc("/", hello)
	http.ListenAndServe(":4000", yamlHandler)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}
