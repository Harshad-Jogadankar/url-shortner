package controller

import (
	"encoding/json"
	"net/http"
	"url-shortner/internal/models"
	"url-shortner/internal/service"
)

type Controller struct {
	svc *service.Service
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (c *Controller) HandleShortener(w http.ResponseWriter, r *http.Request) {
	var req models.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	shortenedUrl, err := c.svc.Shorten(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp models.Response
	resp.ShortenedUrl = shortenedUrl

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func (c *Controller) HandleResolve(w http.ResponseWriter, r *http.Request) {
	var req models.ResolveRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	originalUrl, err := c.svc.Resolve(req.ShortenedUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp models.ResolveResponse
	resp.OriginalUrl = originalUrl

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(resp)
}
