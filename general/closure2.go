package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type appContext struct {
	db      string
	session string
}

//type appHandler struct {
//	handler func(w http.ResponseWriter, r *http.Request) (int, error)
//	c       *appContext
//}
//
//func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	status, err := ah.handler(w, r)
//	if err != nil {
//		log.Printf("HTTP %d: %q", status, err)
//		switch status {
//		case http.StatusNotFound:
//			http.NotFound(w, r)
//		case http.StatusInternalServerError:
//			http.Error(w, http.StatusText(status), status)
//		default:
//			http.Error(w, http.StatusText(status), status)
//		}
//	}
//}

func (app *appContext) IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "db is %q and session is %q\n", app.db, app.session)
	})
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		log.Printf("handing %s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %s", r.Method, r.URL, time.Since(began))
	})
}

func main() {
	context := &appContext{db: "db string", session: "session string"}
	http.Handle("/hello", logHandler(context.IndexHandler()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
