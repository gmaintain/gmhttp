package gmhttp

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewEngine(t *testing.T) {
	req, err := http.NewRequest("GET", "/logger", nil)
	var logbuf bytes.Buffer

	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	testlog := log.Logger{}
	testlog.SetOutput(&logbuf)

	engine := NewEngine(testlog)


	t.Run("log test", func(t *testing.T) {
		want := "hi logger\n"
		err = engine.Get("/logger", func(w http.ResponseWriter, r *http.Request) {
			engine.Logger.Print(want)
		})
		if err != nil {
			t.Fatal(err)
		}
		engine.ServeHTTP(resp, req)
		got := logbuf.String()
		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
	t.Run("conflict register", func(t *testing.T) {
		err1 := engine.Get("/aaa", func(w http.ResponseWriter, r *http.Request) {

		})
		err2 := engine.Post("/aaa", func(w http.ResponseWriter, r *http.Request) {

		})
		err3 := engine.Get("/aaa", func(w http.ResponseWriter, r *http.Request) {

		})
		if err1 != nil || err2 != nil {
			t.Error(err1, err2)
		}

		if err3.Error() != "router has register" {
			t.Error(err)
		}
	})
}

func TestResp(t *testing.T) {
	req, _ := http.NewRequest("POST", "/resp/check", nil)

	engine := NewEngine(log.Logger{})
	engine.Post("/resp/check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(300)
		w.Write([]byte("resp ok!"))
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
