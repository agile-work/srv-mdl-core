package main

import (
	"github.com/agile-work/srv-mdl-core/routes"
	module "github.com/agile-work/srv-mdl-shared"
)

func main() {
	module.ListenAndServe(":3010", routes.Routes())
}
