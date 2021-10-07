package MiniGin

import (
	"net/http"
)
type Engine struct {
	route *Route
}

type handleFun func(c *Context)
var defaultWeb *Engine

func GetNewWeb()*Engine{
	return &Engine{route: getRoute()}
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context:=GetContext(w,req)
	//讲道理，下面的代码有点傻逼，
	e.route.handleContext(context)
}
func (e *Engine) AddGet(pattern string,handle handleFun)  {
	e.route.AddGet(pattern,handle)
}
func (e *Engine) AddPost(pattern string,handle handleFun)  {
	e.route.AddPost(pattern,handle)
}
func (e *Engine) addMethod(method string,pattern string,handle handleFun)  {
	e.route.addMethod(method,pattern,handle)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}





