package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_routes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/login", "POST"},
		{"/user/profile", "GET"},
		{"/static/*", "GET"},
	}

	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range registered {
		t.Run(route.route, func(t *testing.T) {
			if !routeExists(route.route, route.method, chiRoutes) {
				t.Errorf("route '%s' with method '%s' not found", route.route, route.method)
			}
		})
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(testRoute, route) && strings.EqualFold(testMethod, method) {
			found = true
		}

		return nil
	})

	return found
}
