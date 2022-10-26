package httpserver

import "net/http"

type Handlers interface {
	GetInfo(w http.ResponseWriter, r *http.Request)
	GetHealthCheck(w http.ResponseWriter, r *http.Request)
}
