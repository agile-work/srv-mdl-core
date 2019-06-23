package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/user"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostUser sends the request to model creating a new user
func PostUser(res http.ResponseWriter, req *http.Request) {
	user := &user.User{}
	resp := response.New()

	if err := resp.Parse(req, user); err != nil {
		resp.NewError("PostUser response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostUser user new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := user.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostUser", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = user
	resp.Render(res, req)
}

// GetAllUsers return all user instances from the model
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	users := &user.Users{}
	if err := users.LoadAll(opt); err != nil {
		resp.NewError("GetAllUsers", err)
		resp.Render(res, req)
		return
	}
	resp.Data = users
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetUser return only one user from the model
func GetUser(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	user := &user.User{Username: chi.URLParam(req, "username")}
	if err := user.Load(); err != nil {
		resp.NewError("GetUser", err)
		resp.Render(res, req)
		return
	}
	resp.Data = user
	resp.Render(res, req)
}

// UpdateUser sends the request to model updating a user
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	user := &user.User{}
	resp := response.New()

	if err := resp.Parse(req, user); err != nil {
		resp.NewError("UpdateUser user new transaction", err)
		resp.Render(res, req)
		return
	}

	user.Username = chi.URLParam(req, "username")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateUser user new transaction", err)
		resp.Render(res, req)
		return
	}

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateUser user new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := user.Update(trs, util.GetBodyColumns(body)); err != nil {
		trs.Rollback()
		resp.NewError("UpdateUser", err)
		resp.Render(res, req)
		return
	}

	trs.Commit()
	resp.Data = user
	resp.Render(res, req)
}

// DeleteUser sends the request to model deleting a user
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteUser user new transaction", err)
		resp.Render(res, req)
		return
	}

	user := &user.User{Username: chi.URLParam(req, "username")}
	if err := user.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteUser", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}

// AddGroupInUser sends the request to service deleting an user
func AddGroupInUser(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// RemoveGroupFromUser sends the request to service deleting an user
func RemoveGroupFromUser(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllUsersByGroup return all user instances by group from the service
func GetAllUsersByGroup(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllPermissionsByUser return all user instances by group from the service
func GetAllPermissionsByUser(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
