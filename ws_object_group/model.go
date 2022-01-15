package ws_object_group

import (
	"encoding/json"
	model "github.com/mdalbrid/models"
	"time"
)

type ObjectGroup interface {
	ToMap() *map[string]interface{}
}

type changedFields map[string]bool

type Struct struct {
	uuid            model.UUID
	explorationUUID model.UUID
	name            string
	icon            string
	sortWeight      int
	authorUUID      model.UUID
	authorName      string
	dateAdd         time.Time
	dateEdit        time.Time
	deleted         bool

	// _changed - флаг изменения объекта
	_changed bool
	// _changedFields - таблица для хранения измененных свойств
	_changedFields changedFields
}

func (i *Struct) ToMap() *map[string]interface{} {
	return &map[string]interface{}{
		"uuid":            i.uuid.String(),
		"explorationUUID": i.explorationUUID.String(),
		"name":            i.name,
		"icon":            i.icon,
		"sortWeight":      i.sortWeight,
		"authorUUID":      i.authorUUID.String(),
		"authorName":      i.authorName,
		"dateAdd":         i.dateAdd,
		"dateEdit":        i.dateEdit,
		"deleted":         i.deleted,
	}
}

func (i Struct) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ToMap())
}
