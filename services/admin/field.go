package admin

import (
	"errors"
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/field"
	"github.com/agile-work/srv-mdl-core/models/lookup"
	mdlShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreateField insert a new field in the database
func CreateField(fld *field.Field) (string, error) {
	def, err := fld.GetDefinition()
	if err != nil {
		return "CreateField processing definitions", err
	}

	if fld.Type == shared.FieldLookup {
		fldLkpDef := def.(*field.LookupDefinition)
		lkp := lookup.Lookup{}
		err := lkp.Load(fldLkpDef.LookupCode)
		if err != nil {
			return "CreateField load lookup", err
		}
		if !lkp.Active {
			return "CreateField load lookup", errors.New("invalid lookup code")
		}

		lkpDef, err := lkp.GetDefinition()
		if err != nil {
			return "CreateField lookup get definition", err
		}

		fldLkpDef.LookupValue, fldLkpDef.LookupLabel = lkpDef.GetValueAndLabel()
		if err != nil {
			return "CreateField lookup get value and label", err
		}

		if fldLkpDef.Type != shared.FieldLookupStatic {
			lkpDynDef := lkpDef.(*lookup.DynamicDefinition)
			for _, p := range lkpDynDef.Params {
				param := field.LookupParam{
					Code:     p.Code,
					DataType: p.DataType,
				}
				fldLkpDef.LookupParams = append(fldLkpDef.LookupParams, param)
			}
		}
	}

	fld.SetDefinition(def)
	fld.Create()
	return "", nil
}

// LoadAllFields returns all fields from a schema
func LoadAllFields(r *http.Request) *mdlShared.Response {
	return &mdlShared.Response{
		Code: http.StatusNotImplemented,
	}
}

// LoadField return an specific field from a schema
func LoadField(r *http.Request) *mdlShared.Response {
	return &mdlShared.Response{
		Code: http.StatusNotImplemented,
	}
}

// UpdateField updates the field attributes in the database
func UpdateField(r *http.Request) *mdlShared.Response {
	return &mdlShared.Response{
		Code: http.StatusNotImplemented,
	}
}

// DeleteField deletes an specific field in the database
func DeleteField(r *http.Request) *mdlShared.Response {
	return &mdlShared.Response{
		Code: http.StatusNotImplemented,
	}
}

// AddFieldValidation include a new validation to a field
func AddFieldValidation(r *http.Request) *mdlShared.Response {
	return &mdlShared.Response{
		Code: http.StatusNotImplemented,
	}
}

// UpdateFieldValidation update the validation attributes
func UpdateFieldValidation(r *http.Request) *mdlShared.Response {
	return &mdlShared.Response{
		Code: http.StatusNotImplemented,
	}
}

// DeleteFieldValidation delete a validation from the database
func DeleteFieldValidation(r *http.Request) *mdlShared.Response {
	return &mdlShared.Response{
		Code: http.StatusNotImplemented,
	}
}
