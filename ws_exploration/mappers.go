package ws_exploration

import (
	"github.com/jackc/pgx/v4"
	model "github.com/mdalbrid/models"
)

// language=PostgreSQL prefix="SELECT " suffix="  FROM ws_exploration"
const fullCols = `"uuid", "name", "comment", "tags", "accessType", "views", "authorUUID", "authorName", "dateAdd", "dateEdit", "deleted"`

// ScanPgxRow - заполняет все поля объекта из полученного от БД ответа
func ScanPgxRow(res pgx.Row) (f Exploration, err error) {
	exp := &Struct{}
	var guid, author string
	err = res.Scan(&guid, &exp.name, &exp.comment, &exp.tags, &exp.accessType, &exp.views, &author, &exp.authorName, &exp.dateAdd, &exp.dateEdit, &exp.deleted)
	if err != nil {
		return nil, err
	}

	exp.uuid = model.UUID(guid)
	exp.authorUUID = model.UUID(author)
	exp._changedFields = make(changedFields)
	return exp, err
}

// ScanPgxRows - создает список объектов типа файл по результатам полученным из БД
func ScanPgxRows(rows pgx.Rows) ([]Exploration, error) {
	f := []Exploration{}
	for rows.Next() {
		if _file, e := ScanPgxRow(rows); e == nil {
			f = append(f, _file)
		} else {
			return nil, e
		}
	}
	return f, nil
}
