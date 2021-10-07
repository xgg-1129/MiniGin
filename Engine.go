package MiniGin

import (
	"net/http"
)

type Engine struct {
	route *Route
}

type handleFun func(w http.ResponseWriter,req *http.Request)
var defaultWeb *Engine

func GetNewWeb()*Engine{
	return &Engine{route: getRoute()}
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key:=req.Method+"-"+req.URL.Path
	handle := e.route.handleFunMap[key]
	if handle == nil{
		w.WriteHeader(404)
	}else{
		handle(w,req)
	}
}

func (e *Engine) AddGet(pattern string,handle handleFun)  {
	e.addMethod("GET",pattern,handle)
}
func (e *Engine) AddPost(pattern string,handle handleFun)  {
	e.addMethod("POST",pattern,handle)
}
func (e *Engine) addMethod(method string,pattern string,handle handleFun)  {
		key:=method+"-"+pattern
		e.route.handleFunMap[key]=handle
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}


