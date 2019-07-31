package main

import (
	"github.com/agile-work/srv-mdl-core/routes"
	shared "github.com/agile-work/srv-mdl-shared"
)

// swagg-doc:api
// openapi: 3.0.0
// info:
//   title: Horizon API - Core
//   description: Documentation for the module Core with all endpoits available.
//   version: 0.0.1
// servers:
//   - url: http://localhost:8080/api/v1
//     description: Development server
//
// securitySchemes:
//   ApiKeyAuth:
//     type: apiKey
//     in: header
//     name: Authentication
//
// parameters:
//   language-code:
//     name: language-code
//     in: header
//     description: Used by translation fields to get the data translatate. Use 'all' to return data from all languages.
//     required: true
//     schema:
//       type: string
//       example: pt-br

func main() {
	shared.ListenAndServe("core", "localhost", 3010, installModule, routes.Setup())
}

func installModule(moduleID string) error {
	return nil
}
