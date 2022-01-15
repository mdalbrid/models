package ws_exploration

import (
	"encoding/json"
	model "github.com/mdalbrid/models"
	"time"
)

type Exploration interface {
	ToMap() *map[string]interface{}
}

type changedFields map[string]bool

type Struct struct {
	uuid       model.UUID
	name       string
	comment    string
	tags       []string
	accessType string
	views      int64
	authorUUID model.UUID
	authorName string
	dateAdd    time.Time
	dateEdit   time.Time
	deleted    bool

	// _changed - флаг изменения объекта
	_changed bool
	// _changedFields - таблица для хранения измененных свойств
	_changedFields changedFields
}

func (i *Struct) ToMap() *map[string]interface{} {
	return &map[string]interface{}{
		"uuid":       i.uuid.String(),
		"name":       i.name,
		"comment":    i.comment,
		"tags":       i.tags,
		"accessType": i.accessType,
		"views":      i.views,
		"authorUUID": i.authorUUID.String(),
		"authorName": i.authorName,
		"dateAdd":    i.dateAdd,
		"dateEdit":   i.dateEdit,
		"deleted":    i.deleted,
	}
}

func (i Struct) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.ToMap())
}
