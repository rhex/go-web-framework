package context

import ()

type Context struct{}

func Init() *Context {
	return &Context{}
}
