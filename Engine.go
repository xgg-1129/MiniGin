package MiniGin

import (
	"net/http"
	"strings"
)
type Engine struct {
	route *Route
	*Group
	Groups []*Group
}

type HandleFun func(c *Context)
var defaultWeb *Engine

func GetNewWeb()*Engine{
	engine:=&Engine{
		route: NewRoute(),
	}
	engine.Group=&Group{
		engine: engine,
	}
	engine.Groups=[]*Group{engine.Group}
	return engine
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context:=GetContext(w,req)
	for _,group := range e.Groups{
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			context.handles = append(context.handles, group.middles...)
		}
	}
	e.route.handleContext(context)
}
func (e *Engine) AddGet(pattern string,handle HandleFun)  {
	e.route.AddGet(pattern,handle)
}
func (e *Engine) AddPost(pattern string,handle HandleFun)  {
	e.route.AddPost(pattern,handle)
}
func (e *Engine) addMethod(method string,pattern string,handle HandleFun)  {
	e.route.addMethod(method,pattern,handle)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}


