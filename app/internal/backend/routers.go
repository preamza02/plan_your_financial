package backend

import (
	"net/http"
	"plan_your_financial/internal/backend/hello"
)

type Route struct {
	path    string
	handler http.HandlerFunc
}

func initRouters() []Route {
	return []Route{
		{path: "GET /api/hello", handler: hello.GetHelloHandler},
	}
}
