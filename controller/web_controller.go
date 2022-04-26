package controller

import (
	"encoding/json"
	"net/http"
	"webscraper/config"
	"webscraper/model"
	"webscraper/service"
)

type WebController struct {
	c              *config.Config
	webPageService *service.WebPageService
	IndexPage      []byte
}

func NewWebController(c *config.Config, ws *service.WebPageService) *WebController {
	return &WebController{c: c, webPageService: ws}
}

func (h *WebController) Index(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write(h.IndexPage)
}

func (h *WebController) Analyze(w http.ResponseWriter, r *http.Request) {
	var url model.Request
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result, err := h.webPageService.Analyze(url.URL); err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"result":  result,
		})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

}
