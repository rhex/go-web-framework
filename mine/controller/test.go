package controller

import (
	"errors"
	"github.com/rhex/go-web-framework/mine/context"
	"log"
	"net/http"
)

// TODO when to print log

func PanicHandler(ctx *context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Println("Panic")
	panic("here is panic")
	return nil
}

func ErrorHandler(ctx *context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Println("error")
	return errors.New("this is a error from error handler")
}
