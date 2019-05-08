package main

import (
	"github.com/agile-work/srv-mdl-core/routes"
	moduleShared "github.com/agile-work/srv-mdl-shared"
)

func main() {
	moduleShared.ListenAndServe(":3010", routes.Setup())
}
