package MiniGin

type Route struct {
	handleFunMap map[string]handleFun
}

func getRoute() *Route{
	return &Route{handleFunMap: make(map[string]handleFun)}
}

func (r *Route) AddGet(pattern string,handle handleFun)  {
	r.addMethod("GET",pattern,handle)
}
func (r *Route) AddPost(pattern string,handle handleFun)  {
	r.addMethod("POST",pattern,handle)
}
func (r *Route) addMethod(method string,pattern string,handle handleFun)  {
	key:=method+"-"+pattern
	r.handleFunMap[key]=handle
}

