package backend

import "net/http"

func InitBackend(mux *http.ServeMux) {
	Routers := initRouters()
	for _, route := range Routers {
		mux.HandleFunc(route.path, route.handler)
	}
}
