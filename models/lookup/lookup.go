package lookup

import (
	"encoding/json"
	"fmt"
	"time"

	sharedModels "github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
)

type Definition interface {
	parse(payload json.RawMessage) error
	GetValueAndLabel() (string, string)
}

// Lookup defines the struct of this object
type Lookup struct {
	ID          string                   `json:"id" sql:"id" pk:"true"`
	Code        string                   `json:"code" sql:"code" updatable:"false"`
	Type        string                   `json:"type" sql:"type" updatable:"false"`
	Name        sharedModels.Translation `json:"name" sql:"name" field:"jsonb"`
	Description sharedModels.Translation `json:"description" sql:"description" field:"jsonb"`
	Definitions json.RawMessage          `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false"`
	Active      bool                     `json:"active" sql:"active"`
	CreatedBy   string                   `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time                `json:"created_at" sql:"created_at"`
	UpdatedBy   string                   `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time                `json:"updated_at" sql:"updated_at"`
}

func (l *Lookup) Load(lookupCode string) error {
	lookupCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreLookups)
	condition := builder.Equal(lookupCodeColumn, lookupCode)
	if err := sql.SelectStruct(shared.TableCoreLookups, l, condition); err != nil {
		return err
	}
	return nil
}

func (l *Lookup) GetDefinition() (Definition, error) {
	switch l.Type {
	case shared.LookupStatic:
		def := &StaticDefinition{}
		err := def.parse(l.Definitions)
		return def, err
	case shared.LookupDynamic:
		def := &DynamicDefinition{}
		err := def.parse(l.Definitions)
		return def, err
	}
	return nil, nil
}
