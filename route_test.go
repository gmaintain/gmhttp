package gmhttp

import (
	"fmt"
	"reflect"
	"testing"
)

func NewTestRouter() *router {
	r := NewRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/do/:name", nil)
	r.addRouter("GET", "/hello/*file", nil)
	r.addRouter("POST", "/hello/do", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	NewTestRouter()
	ok := reflect.DeepEqual(parsePattern("/p/:id"), []string{"p", ":id"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/:fileName"), []string{"p", ":fileName"})
	if !ok {
		t.Error("parsePattern error")
	}
}

func TestGetRoutes(t *testing.T) {
	r := NewTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Printf("%d %#v\n", i, n)
	}
}
