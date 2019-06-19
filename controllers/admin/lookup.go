package admin

import (
	"net/http"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// PostLookup sends the request to service creating a new lookup
func PostLookup(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllLookups return all lookup instances from the service
func GetAllLookups(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetLookup return only one lookup from the service
func GetLookup(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateLookup sends the request to service updating a lookup
func UpdateLookup(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteLookup sends the request to service deleting a lookup
func DeleteLookup(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// AddLookupOption add a new option to a lookup
func AddLookupOption(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateLookupOption change lookup option data
func UpdateLookupOption(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteLookupOption delete an option
func DeleteLookupOption(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateLookupOrder delete an option
func UpdateLookupOrder(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateLookupQuery change dynamic lookup query
func UpdateLookupQuery(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateLookupDynamicField change dynamic lookup field
func UpdateLookupDynamicField(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateLookupDynamicParam change dynamic lookup param
func UpdateLookupDynamicParam(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// // PostLookup sends the request to service creating a new lookup
// func PostLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.CreateLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // GetAllLookups return all lookup instances from the service
// func GetAllLookups(res http.ResponseWriter, req *http.Request) {
// 	response := services.LoadAllLookups(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // GetLookup return only one lookup from the service
// func GetLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.LoadLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookup sends the request to service updating a lookup
// func UpdateLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // DeleteLookup sends the request to service deleting a lookup
// func DeleteLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.DeleteLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // AddLookupOption add a new option to a lookup
// func AddLookupOption(res http.ResponseWriter, req *http.Request) {
// 	response := services.AddLookupOption(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupOption change lookup option data
// func UpdateLookupOption(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupOption(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // DeleteLookupOption delete an option
// func DeleteLookupOption(res http.ResponseWriter, req *http.Request) {
// 	response := services.DeleteLookupOption(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupOrder delete an option
// func UpdateLookupOrder(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupOrder(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupQuery change dynamic lookup query
// func UpdateLookupQuery(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupQuery(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupDynamicField change dynamic lookup field
// func UpdateLookupDynamicField(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupDynamicParam(r, "field")

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupDynamicParam change dynamic lookup param
// func UpdateLookupDynamicParam(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupDynamicParam(r, "param")

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }
