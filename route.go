package gmhttp

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]handlerFunc
}

func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]handlerFunc),
	}
}

// 解析相对路由path为数组
// path数组丢给trie来insert
func parsePattern(pattern string) []string {
	parts := make([]string, 0)
	vs := strings.Split(pattern, "/")

	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part == "*" {
				break
			}
		}
	}
	return parts
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.Search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) addRouter(method, pattern string, handler handlerFunc) {
	// 构建树, 构建handler map
	parts := parsePattern(pattern)
	key := method + "-" + pattern

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) handler(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + c.Path
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%s\n", c.Path)
	}

}