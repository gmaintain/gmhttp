package gmhttp

import (
	"fmt"
	"testing"
)

func NewTestRouter() *router {
	r := NewRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	return r
}

func TestGetRoutes(t *testing.T) {
	r := NewTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Printf("%d %#v\n", i, n)
	}

}
