package gmhttp

import (
	"fmt"
	"log"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type engine struct {
	Logger log.Logger
	router map[string]handlerFunc
}

func NewEngine(logger log.Logger) *engine {
	return &engine{Logger: logger, router:make(map[string]handlerFunc)}
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pattern := r.Method + "_" + r.URL.Path
	if fun, ok := e.router[pattern]; ok {
		fun(w, r)
	}else {
		w.WriteHeader(404)
	}
}

func (e *engine) addRouter(r string, h handlerFunc) error  {
	if _, ok := e.router[r]; !ok {
		e.router[r] = h
		return nil
	}
	return fmt.Errorf("router has register")
}

//注册pattern以及执行方法
func (e *engine) Get(pattern string, engineFunc handlerFunc) error {
	pattern = "GET_" + pattern
	return e.addRouter(pattern, engineFunc)
}

func (e *engine) Post(pattern string, engineFunc handlerFunc) error {
	pattern = "POST_" + pattern
	return e.addRouter(pattern, engineFunc)
}
func (e *engine) Put(pattern string, engineFunc handlerFunc) error {
	pattern = "PUT_" + pattern
	return e.addRouter(pattern, engineFunc)
}

func (e *engine) Delete(pattern string, engineFunc handlerFunc) error {
	pattern = "DELETE_" + pattern
	return e.addRouter(pattern, engineFunc)
}

func (e *engine) Options(pattern string, engineFunc handlerFunc) error {
	pattern = "OPTIONS_" + pattern
	return e.addRouter(pattern, engineFunc)
}


func Run() {
	engine := NewEngine(log.Logger{})
	err := http.ListenAndServe(":8080", engine)
	if err != nil {
		engine.Logger.Fatalln(err)
	}
}
