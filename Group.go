package MiniGin

type Group struct {
	engine *Engine //这里是为了通过engine访问route，如果写route，初始化的时候还得想办法和engine的route保持一致
	prefix string
	parent *Group

	middles []HandleFun
}

func (g *Group)NewGroup(prefix string)*Group{
	n := &Group{
		engine: g.engine,
		prefix: prefix,
		parent: g,
	}
	g.engine.Groups=append(g.engine.Groups, n)
	return n
}

func (g *Group) addRoute(method string,pattern string,fun HandleFun)  {
	g.engine.addMethod(method,g.prefix+pattern,fun)
}
func (g *Group) AddGet(pattern string,fun HandleFun)  {
	g.addRoute("GET",pattern,fun)
}
func (g *Group) AddPost(pattern string,fun HandleFun)  {
	g.addRoute("POST",pattern,fun)
}
func (g *Group)RegisterMiddle(middleFun HandleFun){
	g.middles=append(g.middles,middleFun)
}
