package auth

import (
	"net/http"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// Login endpoint to get user credentials and return token
func Login(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
