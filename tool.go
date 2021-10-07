package MiniGin

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//tool包里是一些现成的中间件

func RecoverPanic()HandleFun{
	return func(c *Context) {
		defer func() {
			if err := recover();err!=nil{
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.DoAllNext()
	}
}
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller
	var str strings.Builder
	str.WriteString(message)
	for _,item := range pcs[0:n]{
		pc := runtime.FuncForPC(item)
		file, line := pc.FileLine(item)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}