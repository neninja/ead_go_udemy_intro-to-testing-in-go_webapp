package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application(t *testing.T) {
	var tt = []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}

		t.Log(ip)
	})

	for _, e := range tt {
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	// get context
	ctx := context.Background()

	// put something on context
	ctx = context.WithValue(ctx, contextUserKey, "whatever")

	// call the test func
	ip := app.ipFromContext(ctx)

	// assert
	if !strings.EqualFold("whatever", ip) {
		t.Error("wrong value returned from context")
	}
}
