package main

import (
	"github.com/agile-work/srv-mdl-core/routes"
	mdlShared "github.com/agile-work/srv-mdl-shared"
)

func main() {
	mdlShared.ListenAndServe("Core", "localhost", 3010, routes.Setup())
}
