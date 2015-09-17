http://stackoverflow.com/questions/24714624/passing-context-to-interface-methods
https://www.reddit.com/r/golang/comments/2apjja/passing_context_to_interface_methods/cixpku8
Very good:https://www.nicolasmerouze.com/share-values-between-middlewares-context-golang/

func (c *Context) Close() {
    c.Database.Session.Close()
}

type handler func(http.ResponseWriter, *http.Request, *Context) error

## what the mux name is used for

func sign(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
  //...

  url, err := router.GetRoute("index").URL()
  if err != nil {
      return
  }

  http.Redirect(w, req, url, http.StatusTemporaryRedirect)
  return
}
