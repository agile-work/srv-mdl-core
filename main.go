package main

import (
	"github.com/agile-work/srv-mdl-core/routes"
	mdlShared "github.com/agile-work/srv-mdl-shared"
)

func main() {
	mdlShared.ListenAndServe("core", "localhost", 3010, install, routes.Setup())
}

func install() error {
	return nil
}
