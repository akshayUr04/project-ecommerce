//go:generate swagger generate spec
package main

import (
	"log"

	_ "github.com/akshayur04/project-ecommerce/cmd/api/docs"
	_ "github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	_ "github.com/akshayur04/project-ecommerce/pkg/common/response"

	config "github.com/akshayur04/project-ecommerce/pkg/config"
	di "github.com/akshayur04/project-ecommerce/pkg/di"
)

// @title Ecommerce REST API
// @version 1.0
// @description Ecommerce REST API built using Go Lang, PSQL, REST API following Clean Architecture. Hosted with Ngnix, AWS EC2 and RDS
//
//	Schemes:  https
//
// @contact.name Akshay ur
// @contact.url https://github.com/akshayUr04
// @contact.email akshayur0404@gmail.com
// @license.name MIT
// @host localhost
// @license.url https://opensource.org/licenses/MIT
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
