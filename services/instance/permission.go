package instance

import (
	"fmt"
	"net/http"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	shared "github.com/agile-work/srv-shared"
)

// CreatePermission persists the request body creating a new object in the database
func CreatePermission(r *http.Request) *moduleShared.Response {
	permission := models.EntityInstancePermission{
		InstanceType: chi.URLParam(r, "schema_code"),
		InstanceID:   chi.URLParam(r, "instance_id"),
	}

	return db.Create(r, &permission, "CreatePermission", shared.TableCoreInstancePermissions)
}

// LoadAllPermissions return all instances from the object
func LoadAllPermissions(r *http.Request) *moduleShared.Response {
	// TODO: retornar todos os usuários que podem visualizar pq tem global no schema ou que são da mesma tree
	permissions := []models.EntityInstancePermission{}
	instanceType := chi.URLParam(r, "schema_code")
	instanceTypeColumn := fmt.Sprintf("%s.instance_type", shared.TableCoreInstancePermissions)
	instanceID := chi.URLParam(r, "instance_id")
	instanceIDColumn := fmt.Sprintf("%s.instance_id", shared.TableCoreInstancePermissions)
	condition := builder.And(
		builder.Equal(instanceTypeColumn, instanceType),
		builder.Equal(instanceIDColumn, instanceID),
	)

	return db.Load(r, &permissions, "LoadAllPermissions", shared.TableCoreInstancePermissions, condition)
}

// UpdatePermission return all instances from the object
func UpdatePermission(r *http.Request) *moduleShared.Response {
	permissions := models.EntityInstancePermission{}
	instanceType := chi.URLParam(r, "schema_code")
	instanceTypeColumn := fmt.Sprintf("%s.instance_type", shared.TableCoreInstancePermissions)
	instanceID := chi.URLParam(r, "instance_id")
	instanceIDColumn := fmt.Sprintf("%s.instance_id", shared.TableCoreInstancePermissions)
	permissionID := chi.URLParam(r, "permission_id")
	permissionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreInstancePermissions)
	condition := builder.And(
		builder.Equal(instanceTypeColumn, instanceType),
		builder.Equal(instanceIDColumn, instanceID),
		builder.Equal(permissionIDColumn, permissionID),
	)

	return db.Update(r, &permissions, "UpdatePermission", shared.TableCoreInstancePermissions, condition)
}

// DeletePermission deletes object from the database
func DeletePermission(r *http.Request) *moduleShared.Response {
	instanceType := chi.URLParam(r, "schema_code")
	instanceTypeColumn := fmt.Sprintf("%s.instance_type", shared.TableCoreInstancePermissions)
	instanceID := chi.URLParam(r, "instance_id")
	instanceIDColumn := fmt.Sprintf("%s.instance_id", shared.TableCoreInstancePermissions)
	permissionID := chi.URLParam(r, "permission_id")
	permissionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreInstancePermissions)
	condition := builder.And(
		builder.Equal(instanceTypeColumn, instanceType),
		builder.Equal(instanceIDColumn, instanceID),
		builder.Equal(permissionIDColumn, permissionID),
	)

	return db.Remove(r, "DeletePermission", shared.TableCoreInstancePermissions, condition)
}
