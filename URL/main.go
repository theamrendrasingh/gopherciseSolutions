package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	// "github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// yaml := `
	// 		- path: /urlshort
	// 		url: https://github.com/gophercises/urlshort
	// 		- path: /urlshort-final
	// 		url: https://github.com/gophercises/urlshort/tree/solution
	// 		`

	filename, _ := filepath.Abs("./test.yml")
	yaml, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading yaml file")
		panic(err)
	}

	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
	// http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
