package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{
			"regular",
			"Hello, regular!",
		},
		{
			"cyberlight@golang.org",
			"Hello, gopher cyberlight!",
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf("http://localhost:8080/%s", c.in),
			nil,
		)

		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res := httptest.NewRecorder()
		handler(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("wrong response code: %v", res.Code)
		}

		if !strings.Contains(res.Body.String(), c.out) {
			t.Errorf("wrong response body content: %v", res.Body.String())
		}
	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/cyberlight@golang.org",
			nil,
		)

		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}

		res := httptest.NewRecorder()
		handler(res, req)

		if res.Code != http.StatusOK {
			b.Errorf("wrong response code: %v", res.Code)
		}

		if !strings.Contains(res.Body.String(), "cyberlight") {
			b.Errorf("wrong response body content: %v", res.Body.String())
		}
	}
}
