package controller

import (
	"encoding/json"
	"github.com/rhex/go-web-framework/mine/context"
	"github.com/rhex/go-web-framework/mine/model"
	"log"
	"net/http"
)

// TODO when to print log

func GetTodosHandler(ctx *context.Context, w http.ResponseWriter, r *http.Request) error {
	todos := model.TodosResponse{
		model.TodoResponse{Name: "Write presentation"},
		model.TodoResponse{Name: "Host meetup"},
	}
	s, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(s)
	log.Println("Get Todos")

	return nil
}
