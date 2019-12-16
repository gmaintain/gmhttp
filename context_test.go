package gmhttp

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tlog = log.New(os.Stderr, "", log.LstdFlags)

func TestContext_Status(t *testing.T) {
	req, _ := http.NewRequest("GET", "/aaa", nil)
	resp := httptest.NewRecorder()
	engine := NewEngine(tlog)
	engine.Get("/aaa", func(c *Context) {
		c.Status(205)
	})
	engine.ServeHTTP(resp, req)
	if resp.Code != 205 {
		t.Errorf("want: %v, got: %v", 205, resp.Code)
	}
}
