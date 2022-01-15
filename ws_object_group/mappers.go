package ws_object_group

import (
	"github.com/jackc/pgx/v4"
	model "github.com/mdalbrid/models"
)

// language=PostgreSQL prefix="SELECT " suffix="  FROM ws_object_group"
const fullCols = `"uuid", "explorationUUID", "name", "icon", "sortWeight", "authorUUID", "authorName", "dateAdd", "dateEdit", "deleted"`

// ScanPgxRow - заполняет все поля объекта из полученного от БД ответа
func ScanPgxRow(res pgx.Row) (f ObjectGroup, err error) {
	group := &Struct{}
	var guid, explorationUUID, authorUUID string
	err = res.Scan(&guid, &explorationUUID, &group.name, &group.icon, &group.sortWeight, &authorUUID, &group.authorName, &group.dateAdd, &group.dateEdit, &group.deleted)
	if err != nil {
		return nil, err
	}
	group.uuid = model.UUID(guid)
	group.authorUUID = model.UUID(authorUUID)
	group.explorationUUID = model.UUID(explorationUUID)
	group._changedFields = make(changedFields)
	return group, err
}

// ScanPgxRows - создает список объектов типа файл по результатам полученным из БД
func ScanPgxRows(rows pgx.Rows) ([]ObjectGroup, error) {
	f := []ObjectGroup{}
	for rows.Next() {
		if _file, e := ScanPgxRow(rows); e == nil {
			f = append(f, _file)
		} else {
			return nil, e
		}
	}
	return f, nil
}
