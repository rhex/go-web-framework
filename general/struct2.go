package main

import (
	"fmt"
	"log"
	"net/http"
)

type appContext struct {
	db      string
	session string
}

type appHandler struct {
	handler func(app *appContext, w http.ResponseWriter, r *http.Request) (int, error)
	*appContext
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := ah.handler(ah.appContext, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func IndexHandler(app *appContext, w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Fprintf(w, "db is %q and session is %q\n", app.db, app.session)
	return 200, nil
}

func main() {
	context := &appContext{db: "db string", session: "session string"}
	http.Handle("/hello", appHandler{IndexHandler, context})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
