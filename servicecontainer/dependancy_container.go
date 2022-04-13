package servicecontainer

import (
	"webscraper/config"
	"webscraper/controller"
	"webscraper/service"
)

type DependancyContainer struct {
	config *config.Config

	// controller
	HealthController *controller.HealthController
	WebController    *controller.WebController

	// service
	WebService *service.WebPageService
}

func (d *DependancyContainer) Init() *config.Config {
	d.config = config.LoadAll("config/config.yml")

	d.WebService = service.NewWebPageService(d.config)

	d.HealthController = controller.NewHealthController(d.config)
	d.WebController = controller.NewWebController(d.config, d.WebService)

	return d.config
}
