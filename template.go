package MiniGin

import (
	"net/http"
	"path"
)
/*=========golang的自带库实现了根据路径返回静态资源的方法，这里仅仅需要实现文件映射即可
day6的知识一直是比较薄弱的部分，需要多看多思考
*/
func (g *Group) createStaticHandler(relativePath string,fs http.FileSystem)HandleFun{
	absolutePath:=path.Join(g.prefix,relativePath)
	server := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func (ctx *Context){
		server.ServeHTTP(ctx.W,ctx.Req)
	}
}
//指定一个文件见的内容都是静态文件，并且指定访问它的前缀URL
func (g *Group) Static(prefix string,fileRoot string)  {
	//任何访问prefix开头的路由都会转到这里
	urlpattern := path.Join(prefix,"/*filename")
	//获取处理这些路由的方法
	handle := g.createStaticHandler(prefix, http.Dir(fileRoot))
	//添加到路由表
	g.AddGet(urlpattern,handle)
}
/*===============================================================*/
