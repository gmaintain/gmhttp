package gmhttp

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewEngine(t *testing.T) {
	req, err := http.NewRequest("GET", "/logger", nil)
	var logbuf bytes.Buffer

	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	tlog.SetOutput(&logbuf)
	engine := NewEngine(tlog)

	t.Run("log test", func(t *testing.T) {
		want := "hi logger\n"
		err = engine.Get("/logger", func(c *Context) {
			engine.Logger.Println(want)
		})
		if err != nil {
			t.Fatal(err)
		}
		engine.ServeHTTP(resp, req)
		got := logbuf.String()
		if strings.Contains(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestResp(t *testing.T) {
	req, _ := http.NewRequest("POST", "/resp/check", nil)

	engine := NewEngine(tlog)
	engine.Post("/resp/check", func(c *Context) {
		c.Writer.WriteHeader(300)
		c.Writer.Write([]byte("resp ok!"))
	})
	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	if resp.Code != 300 {
		t.Error("status error")
	}
	if resp.Body.String() != "resp ok!" {
		t.Error("resp body error")
	}
}

func TestContext_Json(t *testing.T) {
	req, _ := http.NewRequest("POST", "/resp/check", nil)

	engine := NewEngine(tlog)
	engine.Post("/resp/check", func(c *Context) {
		c.Json(http.StatusOK, H{"Name": map[string]interface{}{"name": "ss", "age": 12}})
	})
	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Error(t.Name() + "status error")
	}
	if resp.Body.String() != `{"Name":{"age":12,"name":"ss"}}
` {
		t.Error(t.Name()+"resp body error, got:", resp.Body.String())
	}
}
