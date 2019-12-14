package gmhttp

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContext_Status(t *testing.T) {
	req, _ := http.NewRequest("GET", "/aaa", nil)
	resp := httptest.NewRecorder()
	engine := NewEngine(log.Logger{})
	engine.Get("/aaa", func(c *Context) {
		c.Status(201)
	})
	engine.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Error("status func test err")
	}
}
