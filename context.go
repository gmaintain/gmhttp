package gmhttp

import (
	"encoding/json"
	"log"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Logger  log.Logger
}

func NewContext(writer http.ResponseWriter, request *http.Request, logger log.Logger) *Context {
	return &Context{Writer: writer, Request: request, Logger: logger}
}

func (c Context) Json(httpcode int, data interface{}) {
	c.Writer.WriteHeader(httpcode)
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Logger.Fatalln("data 数据结构解析为json异常")
	}
	c.Writer.Write(jsonData)
}
