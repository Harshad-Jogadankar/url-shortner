package router

import (
	"net/http"
	"url-shortner/internal/controller"
)

func NewRouter(c *controller.Controller) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", c.HandleShortener)
	mux.HandleFunc("GET /resolve", c.HandleResolve)
	mux.HandleFunc("GET /{code}", c.HandleRedirect)

	return mux
}
