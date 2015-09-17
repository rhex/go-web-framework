// http://thenewstack.io/make-a-restful-json-api-go/
package main

import (
	//	"html"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	// # Classic way
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
