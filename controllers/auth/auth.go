package auth

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/user"
)

// Login endpoint to get user credentials and return token
func Login(res http.ResponseWriter, req *http.Request) {
	user := &user.User{}
	resp := response.New()

	if err := resp.Parse(req, user); err != nil {
		resp.NewError("Login response load", err)
		resp.Render(res, req)
		return
	}

	if err := user.Login(); err != nil {
		resp.NewError("Login", err)
		resp.Render(res, req)
		return
	}

	resp.Data = user
	resp.Render(res, req)
}
