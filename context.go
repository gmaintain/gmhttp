package gmhttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type H map[string]interface{}
type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Logger     log.Logger
	StatusCode int
	Method     string
	Path       string
	Params     map[string]string
}

func NewContext(writer http.ResponseWriter, request *http.Request, logger log.Logger) *Context {
	return &Context{
		Writer:  writer,
		Request: request,
		Logger:  logger,
		Method:  request.Method,
		Path:    request.URL.Path,
	}
}

func (c *Context) Json(httpcode int, data interface{}) {
	c.Writer.WriteHeader(httpcode)
	//注意这里会产生json之后添加\n换行符
	encode := json.NewEncoder(c.Writer)
	err := encode.Encode(data)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, val string) {
	c.Writer.Header().Set(key, val)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}
