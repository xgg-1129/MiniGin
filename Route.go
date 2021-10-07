package MiniGin

import "fmt"

type Route struct {
	roots map[string]*node
	handleFunMap map[string]HandleFun
}

func NewRoute() *Route{
	return &Route{
		handleFunMap: make(map[string]HandleFun),
		roots: make(map[string]*node),
	}
}

func (r *Route) AddGet(pattern string,handle HandleFun)  {
	r.addMethod("GET",pattern,handle)
}
func (r *Route) AddPost(pattern string,handle HandleFun)  {
	r.addMethod("POST",pattern,handle)
}
func (r *Route) addMethod(method string,pattern string,handle HandleFun)  {
	key:=method+"-"+pattern
	r.handleFunMap[key]=handle
	if r.roots[method]==nil{
		r.roots[method]=&node{
			Pattern:  "",
			Value:    "",
			Dim:      false,
			children: nil,
		}
	}
	r.roots[method].Insert(pattern,ParsePattern(pattern),0)
}
func (r *Route) GetRoute(ctx *Context) *node {
	parts:=ParsePattern(ctx.Req.URL.Path)
	root:=r.roots[ctx.Req.Method]
	if root == nil{
		return nil
	}
	return root.Search(parts,0)
}

func (r *Route) handleContext(ctx *Context) {
	route:= r.GetRoute(ctx)
	if route !=nil{
		key:=ctx.Req.Method+"-"+route.Pattern
		handfun:=r.handleFunMap[key]
		ctx.handles=append(ctx.handles,handfun)
		fmt.Println(ctx.handles)
		ctx.DoAllNext()
	}else{
		ctx.String(404,"file %s not found\n",ctx.Req.URL.Path)
	}
}