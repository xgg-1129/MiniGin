package MiniGin

import (
	"html/template"
	"net/http"
	"strings"
)
type Engine struct {
	route *Route
	*Group
	Groups []*Group
	//一个template里面根据名字可以划分为多个模板
	templates *template.Template
	funcMap   template.FuncMap

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
	context.engine=e
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
//加载整个文件夹为templates模板
/*
htmlTemplates := template.Must(template.New("").ParseGlob("D:\\Environment\\ProjectGo\\src\\Gee\\Template\\*"))
templates := htmlTemplates.Templates()
for _,item := range templates{
	fmt.Println(item.Name())
}*/

//把某个文件夹下的文件都注册到模板中
func (e *Engine) RegisterTemplate(fileSystem string)  {
	e.templates=template.Must(template.New("").Funcs(e.funcMap).ParseGlob(fileSystem))
}
func (e *Engine) RegisterFunMap(funmap template.FuncMap)  {
	e.funcMap=funmap
}


