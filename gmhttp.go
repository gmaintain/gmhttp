package gmhttp

import (
	"log"
	"net/http"
)

type handlerFunc func(c *Context)

type engine struct {
	Logger log.Logger
	router *router
}

func NewEngine(logger log.Logger) *engine {
	return &engine{Logger: logger, router: NewRouter()}
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r, e.Logger)
	e.router.handler(c)
}

func (e *engine) addRouter(method, pattern string, h handlerFunc) error {
	e.router.addRouter(method, pattern, h)
	return nil
}

//注册pattern以及执行方法
func (e *engine) Get(pattern string, engineFunc handlerFunc) error {
	return e.addRouter("GET", pattern, engineFunc)
}

func (e *engine) Post(pattern string, engineFunc handlerFunc) error {
	return e.addRouter("POST", pattern, engineFunc)
}

func (e *engine) Put(pattern string, engineFunc handlerFunc) error {
	return e.addRouter("PUT", pattern, engineFunc)

}

func (e *engine) Delete(pattern string, engineFunc handlerFunc) error {
	return e.addRouter("DELETE", pattern, engineFunc)

}

func (e *engine) Options(pattern string, engineFunc handlerFunc) error {
	return e.addRouter("OPTIONS", pattern, engineFunc)
}

func (e *engine) Run() {
	e.Logger.Fatalln(http.ListenAndServe(":8080", e))
}
