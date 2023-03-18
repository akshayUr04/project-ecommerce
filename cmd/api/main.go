package main

import (
	"log"

	_ "github.com/akshayur04/project-ecommerce/cmd/api/docs"
	_ "github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	_ "github.com/akshayur04/project-ecommerce/pkg/common/response"

	config "github.com/akshayur04/project-ecommerce/pkg/config"
	di "github.com/akshayur04/project-ecommerce/pkg/di"
)

// @title Go + Gin E Commerce API
// @version 1.0
// @description This is an ECommerce server . You can visit the GitHub repository at https://github.com/akshayUr04/project-ecommerce-

// @contact.name API Support
// @contact.url https://github.com/akshayUr04/project-ecommerce-
// @contact.email akshayur0404@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @query.collection.format multi

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
