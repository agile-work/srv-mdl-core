package auth

// swagg-doc:controller
// tags:
//   - name: Auth
//     description: Enpoints to deal with authorization

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/user"
)

// Login swagg-doc:endpoint POST /login
// summary: Send user credential to get a valid token.
// tags:
//   - Auth
// description: Send user credential to get a valid token.
// requestBody:
//   description: User credentials
//   required: true
//   content:
//     application/json:
//       schema:
//         type: object
//         properties:
//           email:
//             type: string
//           password:
//             type: string
//         required:
//           - email
//           - password
// responses:
//   '200':
//     description: The User data with a valid token
//     content:
//       application/json:
//         schema:
//           $ref: '#/components/schemas/User'
//   '400':
//     description: Invalid request body
//   '401':
//     description: Invalid password for this user or the user is not active
//   '404':
//     description: Invalid email user not found
//   '500':
//     description: Problems related to database connection
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
