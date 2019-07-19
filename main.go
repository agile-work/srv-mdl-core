package main

import (
	"github.com/agile-work/srv-mdl-core/routes"
	mdlShared "github.com/agile-work/srv-mdl-shared"
)

func main() {
	mdlShared.ListenAndServe("core", "localhost", 3010, installModule, routes.Setup())
}

func installModule() error {
	return nil
}
