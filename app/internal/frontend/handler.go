package frontend

import "net/http"

func InitFrontend(mux *http.ServeMux) {
	mux.Handle("/",http.FileServer(http.Dir("./internal/frontend/static")))
}
