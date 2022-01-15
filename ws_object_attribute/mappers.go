package ws_object_attribute

import (
	"github.com/jackc/pgx/v4"
	model "github.com/mdalbrid/models"
)

// language=PostgreSQL prefix="SELECT " suffix="  FROM ws_object_groups"
const fullCols = `"uuid", "objectUUID", "name", "value", "sortWeight", "authorUUID", "authorName", "dateAdd", "dateEdit", "deleted"`

// ScanPgxRow - заполняет все поля объекта из полученного от БД ответа
func ScanPgxRow(res pgx.Row) (f ObjectAttribute, err error) {
	group := &Struct{}
	var guid, objectUUID, authorUUID string
	err = res.Scan(&guid, &objectUUID, &group.name, &group.value, &group.sortWeight, &authorUUID, &group.authorName, &group.dateAdd, &group.dateEdit, &group.deleted)
	if err != nil {
		return nil, err
	}
	group.uuid = model.UUID(guid)
	group.authorUUID = model.UUID(authorUUID)
	group.objectUUID = model.UUID(objectUUID)
	group._changedFields = make(changedFields)
	return group, err
}

// ScanPgxRows - создает список объектов типа файл по результатам полученным из БД
func ScanPgxRows(rows pgx.Rows) (f []ObjectAttribute, err error) {
	f = []ObjectAttribute{}
	for rows.Next() {
		if _file, e := ScanPgxRow(rows); e == nil {
			f = append(f, _file)
		} else {
			return nil, err
		}
	}
	return
}
