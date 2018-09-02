// tech_test manipulation REST API
//
// The purpose of this application is to provide
// tech_test manipulation REST API
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"tech_test/rest"
)

var (
	Version   = rest.VERSION_V1
	BuildDate string
	GitCommit string
)

// http://petstore.swagger.io/?url=http%3A%2F%2Flocalhost%3A3000%2Fswagger.json

func main() {
	log.Print("Version: ", Version)
	log.Print("Build Time: ", BuildDate)
	log.Print("GitCommit: ", GitCommit)

	rest.CreateRoutes()
}
