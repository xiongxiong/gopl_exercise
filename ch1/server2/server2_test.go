package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	routes()
}

func TestCount(t *testing.T) {
	t.Log("Test count.")
	{
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal("\terror / - ", err)
		}
		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\t/ Should receive 200", ballotX, rw.Code)
		}
		t.Log("\t/ Should receive 200", checkMark)

		req, err = http.NewRequest("GET", "/count", nil)
		if err != nil {
			t.Fatal("\terror /count - ", err)
		}
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\t/count Should receive 200", ballotX, rw.Code)
		}
		t.Log("\t/count Should receive 200", checkMark)

		res := string(rw.Body.Bytes())
		if !strings.Contains(res, "Count 1") {
			t.Fatalf("\t!!! Should count 1, got -- %v - %v", ballotX, res)
		}
	}
}
