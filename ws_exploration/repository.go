package ws_exploration

import (
	"fmt"
	"github.com/gofrs/uuid"
	model "github.com/mdalbrid/models"
	r "github.com/mdalbrid/models/db"
	"strings"
)

type Filter struct {
	FilterName string
}

type FilterObject struct {
	Offset int
	Limit  int
	Orders []Order
	Filter Filter
}

type Order struct {
	Column    string
	Direction string
}

func List(filter FilterObject) ([]Exploration, error) {
	// language=PostgreSQL
	where := make([]string, 0)
	params := []interface{}{}

	sql := fmt.Sprintf(`SELECT %s FROM ws_exploration WHERE deleted = false`, fullCols)

	orders := make([]string, 0)
	for _, b := range filter.Orders {
		orders = append(orders, fmt.Sprintf(`"%s" %s`, b.Column, b.Direction))
	}
	if len(orders) == 0 {
		// language=PostgreSQL prefix="SELECT * FROM ws_exploration ORDER BY "
		orders = append(orders, `"dateAdd" DESC`)
	}

	where, params = filter.Filter.ApplyToParams(where, params)
	whereString := strings.Join(where, " AND ")
	if len(whereString) > 0 {
		sql = fmt.Sprintf(`%s AND %s`, sql, whereString)
	}

	sql = fmt.Sprintf(`%s ORDER BY %s`, sql, strings.Join(orders, ", "))
	sql = fmt.Sprintf(`%s OFFSET %d`, sql, filter.Offset)
	if filter.Limit > 0 {
		sql = fmt.Sprintf(`%s LIMIT %d`, sql, filter.Limit)
	}

	res, err := r.Pool.Query(r.Ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return ScanPgxRows(res)
}

func Count(filter FilterObject) (int64, error) {
	// language=PostgreSQL
	sql := `SELECT COUNT(*) FROM ws_exploration WHERE deleted = false`
	count := int64(0)
	err := r.Pool.QueryRow(r.Ctx, sql).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func Create(name string, comment string, accessType string, tags []string) (Exploration, error) {
	guid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// language=PostgreSQL
	sql := `INSERT INTO ws_exploration ("uuid", "name", "comment", "tags", "accessType", "authorUUID", "authorName") VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols),
		guid.String(), name, comment, tags, accessType, model.NullUUID.String(), "")
	return ScanPgxRow(row)
}

func Edit(guid model.UUID, name string, comment string, accessType string, tags []string) (Exploration, error) {
	// language=PostgreSQL
	sql := `UPDATE ws_exploration SET ("name", "comment", "tags", "accessType", "dateEdit") = ($2,$3,$4,$5,now()) WHERE "uuid" = $1
RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols),
		guid.String(), name, comment, tags, accessType)
	return ScanPgxRow(row)
}

func Delete(guid model.UUID) (Exploration, error) {
	// language=PostgreSQL
	sql := `UPDATE ws_exploration SET "deleted" = true WHERE "uuid" = $1 RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String())
	return ScanPgxRow(row)
}

func Get(guid model.UUID) (Exploration, error) {
	// language=PostgreSQL
	sql := `SELECT %s FROM ws_exploration WHERE "uuid" = $1 AND deleted = false`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String())
	return ScanPgxRow(row)
}
