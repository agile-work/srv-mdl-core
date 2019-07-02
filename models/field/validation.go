package field

import "github.com/agile-work/srv-shared/sql-builder/db"

// AddFieldValidation include a new validation to a field
func (f *Field) AddFieldValidation(trs *db.Transaction) error {
	return nil
}

// UpdateFieldValidation update the validation attributes
func (f *Field) UpdateFieldValidation(trs *db.Transaction) error {
	return nil
}

// DeleteFieldValidation delete a validation from the database
func (f *Field) DeleteFieldValidation(trs *db.Transaction) error {
	return nil
}
