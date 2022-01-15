package ws_object_attribute

import (
	"fmt"
	"github.com/gofrs/uuid"
	model "github.com/mdalbrid/models"
	r "github.com/mdalbrid/models/db"
	"strings"
)

type Filter struct {
	ObjectUUID model.UUID
}

type FilterObject struct {
	Offset int
	Limit  int
	Filter Filter
	Orders []Order
}

type Order struct {
	Column    string
	Direction string
}

func List(filter FilterObject) ([]ObjectAttribute, error) {
	// language=PostgreSQL
	sql := fmt.Sprintf(`SELECT %s FROM ws_object_attribute WHERE "deleted" = false AND "objectUUID" = $1`, fullCols)

	orders := make([]string, 0)
	for _, b := range filter.Orders {
		orders = append(orders, fmt.Sprintf(`"%s" %s`, b.Column, b.Direction))
	}
	if len(orders) == 0 {
		// language=PostgreSQL prefix="SELECT * FROM ws_object_attribute ORDER BY "
		orders = append(orders, `"dateAdd" DESC`)
	}

	sql = fmt.Sprintf(`%s ORDER BY %s`, sql, strings.Join(orders, ", "))
	sql = fmt.Sprintf(`%s OFFSET %d`, sql, filter.Offset)
	if filter.Limit > 0 {
		sql = fmt.Sprintf(`%s LIMIT %d`, sql, filter.Limit)
	}

	res, err := r.Pool.Query(r.Ctx, sql, filter.Filter.ObjectUUID.String())
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return ScanPgxRows(res)
}

func Count(filter FilterObject) (int64, error) {
	// language=PostgreSQL
	sql := `SELECT COUNT(*) FROM ws_object_attribute WHERE deleted = false AND "objectUUID" = $1`
	count := int64(0)
	err := r.Pool.QueryRow(r.Ctx, sql, filter.Filter.ObjectUUID.String()).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func Create(objectUUID model.UUID, name string, icon string) (ObjectAttribute, error) {
	guid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// language=PostgreSQL
	sql := `INSERT INTO ws_object_attribute ("uuid", "objectUUID", "name", "value", "authorUUID", "authorName") VALUES ($1,$2,$3,$4,$5,$6)
RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols),
		guid.String(), objectUUID.String(), name, icon, model.NullUUID.String(), "")
	return ScanPgxRow(row)
}

func Edit(guid model.UUID, name string, icon string, sortWeight int) (ObjectAttribute, error) {
	// language=PostgreSQL
	sql := `UPDATE ws_object_attribute SET ("name", "icon", "sortWeight", "dateEdit") = ($2,$3,$4,now()) WHERE "uuid" = $1 RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String(), name, icon, sortWeight)
	return ScanPgxRow(row)
}

func Delete(guid model.UUID) (ObjectAttribute, error) {
	// language=PostgreSQL
	sql := `UPDATE ws_object_attribute SET "deleted" = true WHERE "uuid" = $1 RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String())
	return ScanPgxRow(row)
}

func Get(guid model.UUID) (ObjectAttribute, error) {
	// language=PostgreSQL
	sql := `SELECT %s FROM ws_object_attribute WHERE "uuid" = $1 AND deleted = false`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String())
	return ScanPgxRow(row)
}
