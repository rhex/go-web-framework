package main

import (
	"flag"
	//	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/justinas/alice"
	"github.com/rhex/go-web-framework/mine/context"
	"github.com/rhex/go-web-framework/mine/controller"
	"github.com/rhex/go-web-framework/mine/middleware"
)

const (
	serviceName = "todoApp"
)

func main() {
	var port string
	var err error
	// TODO: how to use flag
	flag.StringVar(&port, "port", "8082", "Listening port, default 8082")
	flag.Parse()

	ctx := context.Init()

	commonHandlers := alice.New(middleware.Logger, middleware.Panic)

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/todos", commonHandlers.Append(middleware.Accept).Then(wrap(ctx, controller.GetTodosHandler))).Methods("GET")
	router.Handle("/panic", commonHandlers.Then(wrap(ctx, controller.PanicHandler))).Methods("GET")
	router.Handle("/error", commonHandlers.Then(wrap(ctx, controller.ErrorHandler))).Methods("GET")

	log.Println("Starting Server on", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Failed to start, Error: %s", err)
	}

}

type handler func(ctx *context.Context, w http.ResponseWriter, r *http.Request) error

func wrap(ctx *context.Context, h handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h(ctx, w, r); err != nil {
			log.Println("err in http handler")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
