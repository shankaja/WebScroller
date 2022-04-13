package route

import (
	"fmt"
	"log"
	"net/http"
	"webscraper/config"
	"webscraper/controller"

	"github.com/gorilla/mux"
)

type Router struct {
	healthController *controller.HealthController
	webController    *controller.WebController
	c                *config.Config
}

func NewRouter(hc *controller.HealthController, wc *controller.WebController, c *config.Config) *Router {
	return &Router{healthController: hc, c: c, webController: wc}
}

func (t *Router) InitializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/health", t.healthController.GetHealth)
	r.HandleFunc("/index", t.webController.Index)
	r.HandleFunc("/analyze", t.webController.Analyze)

	host := fmt.Sprintf("%s:%s", t.c.Service.URL, t.c.Service.Port)
	log.Fatal(http.ListenAndServe(host, r))
}
