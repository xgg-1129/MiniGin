package MiniGin

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type H map[string]interface{}
type Context struct {
	W http.ResponseWriter
	Req *http.Request

	index int  //context的当前执行函数

	//handle里面的0:len是中间件函数，最后一个是注册的handle函数
	handles []HandleFun

	engine *Engine
}

func GetContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		W:   w,
		Req: req,
		index: -1,
		handles: make([]HandleFun,0),
	}
}

func (c *Context)SetStatus(statuscode int)  {
	c.W.WriteHeader(statuscode)
}
func (c *Context) SetHeader(key,value string)  {
	c.W.Header().Set(key,value)
}
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.W.Write(data)
}


func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.W.Write([]byte(html))
}
func (c *Context) HtmlTemplate(code int, name string,data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	templates:=c.engine.templates.Templates()
	fmt.Println(templates)
	if err := c.engine.templates.ExecuteTemplate(c.W, name, data);err !=nil{
		fmt.Fprint(c.W,err.Error())
	}
}

/*=============================中间件==============================*/
func (c *Context) DoAllNext()  {
	c.index++
	n:=len(c.handles)
	for ;c.index<n;c.index++{
		c.handles[c.index](c)
	}
}


