package ws_object

import (
	"encoding/json"
	model "github.com/mdalbrid/models"
	"time"
)

type Object interface {
	ToMap() *map[string]interface{}
}

type changedFields map[string]bool

type Struct struct {
	uuid            model.UUID
	explorationUUID model.UUID
	groupUUID       model.UUID
	name            string
	image           string
	country         string
	comment         string
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
		"groupUUID":       i.groupUUID.String(),
		"name":            i.name,
		"image":           i.image,
		"country":         i.country,
		"comment":         i.comment,
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
