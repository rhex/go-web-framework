// http://thenewstack.io/make-a-restful-json-api-go/
package main

import (
	"fmt"
	"log"
	"net/http"
)

type HelloHandler struct {
	db string
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var name string
	name = "hello"
	fmt.Fprintf(w, "hi %s!\n", name)
}

func main() {
	http.Handle("/hello", &HelloHandler{db: "db string"})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
