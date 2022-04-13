package main

import (
	_ "embed"
	"webscraper/route"
	"webscraper/servicecontainer"
)

//go:embed ui/index.html
var indexPage []byte

func main() {

	dc := &servicecontainer.DependancyContainer{}
	cfg := dc.Init()
	dc.WebController.IndexPage = indexPage

	route.NewRouter(dc.HealthController, dc.WebController, cfg).InitializeRouter()
}
