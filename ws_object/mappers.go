package ws_object

import (
	"github.com/jackc/pgx/v4"
	model "github.com/mdalbrid/models"
)

// language=PostgreSQL prefix="SELECT " suffix="  FROM ws_object"
const fullCols = `"uuid", "explorationUUID", "groupUUID", "name", "image", "country", "comment", "sortWeight", "authorUUID", "authorName", "dateAdd", "dateEdit", "deleted"`

// ScanPgxRow - заполняет все поля объекта из полученного от БД ответа
func ScanPgxRow(res pgx.Row) (f Object, err error) {
	object := &Struct{}
	var guid, explorationUUID, groupUUID, authorUUID string
	err = res.Scan(&guid, &explorationUUID, &groupUUID, &object.name, &object.image, &object.country, &object.comment, &object.sortWeight, &authorUUID, &object.authorName, &object.dateAdd, &object.dateEdit, &object.deleted)
	if err != nil {
		return nil, err
	}
	object.uuid = model.UUID(guid)
	object.authorUUID = model.UUID(authorUUID)
	object.explorationUUID = model.UUID(explorationUUID)
	object.groupUUID = model.UUID(groupUUID)
	object._changedFields = make(changedFields)
	return object, err
}

// ScanPgxRows - создает список объектов типа файл по результатам полученным из БД
func ScanPgxRows(rows pgx.Rows) ([]Object, error) {
	f := []Object{}
	for rows.Next() {
		if _file, e := ScanPgxRow(rows); e == nil {
			f = append(f, _file)
		} else {
			return nil, e
		}
	}
	return f, nil
}
