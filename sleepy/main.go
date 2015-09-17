package main

import (
	"github.com/rhex/go-web-framework/sleepy/sleepy"
	"net/http"
	"net/url"
)

type Item struct{}
type Single struct{}

func (item Item) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	items := []string{"item1", "item2"}
	data := map[string][]string{"items": items}
	return 200, data, http.Header{"Content-type": {"application/json"}}
}

func (item Item) Put(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	items := []string{"item1"}
	data := map[string][]string{"items": items}
	return 200, data, http.Header{"Content-type": {"application/json"}}
}

func (single Single) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	items := []string{"item1", "item2", "this is single"}
	data := map[string][]string{"items": items}
	return 200, data, http.Header{"Content-type": {"application/json"}}
}

func main() {
	item := new(Item)
	single := new(Single)

	api := sleepy.NewAPI()
	api.AddResource(item, "/items")
	api.AddResource(single, "/items/123")
	api.Start(3000)
}
