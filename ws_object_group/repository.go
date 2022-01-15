package ws_object_group

import (
	"fmt"
	"github.com/gofrs/uuid"
	model "github.com/mdalbrid/models"
	r "github.com/mdalbrid/models/db"
	"strings"
)

type Filter struct {
	ExplorationUUID model.UUID
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

func List(filter FilterObject) ([]ObjectGroup, error) {
	// language=PostgreSQL
	sql := fmt.Sprintf(`SELECT %s FROM ws_object_groups WHERE "deleted" = false AND "explorationUUID" = $1`, fullCols)

	orders := make([]string, 0)
	for _, b := range filter.Orders {
		orders = append(orders, fmt.Sprintf(`"%s" %s`, b.Column, b.Direction))
	}
	if len(orders) == 0 {
		// language=PostgreSQL prefix="SELECT * FROM ws_object_groups ORDER BY "
		orders = append(orders, `"dateAdd" DESC`)
	}

	sql = fmt.Sprintf(`%s ORDER BY %s`, sql, strings.Join(orders, ", "))
	sql = fmt.Sprintf(`%s OFFSET %d`, sql, filter.Offset)
	if filter.Limit > 0 {
		sql = fmt.Sprintf(`%s LIMIT %d`, sql, filter.Limit)
	}

	res, err := r.Pool.Query(r.Ctx, sql, filter.Filter.ExplorationUUID.String())
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return ScanPgxRows(res)
}

func Count(filter FilterObject) (int64, error) {
	// language=PostgreSQL
	sql := `SELECT COUNT(*) FROM ws_object_groups WHERE deleted = false AND "explorationUUID" = $1`
	count := int64(0)
	err := r.Pool.QueryRow(r.Ctx, sql, filter.Filter.ExplorationUUID.String()).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func Create(explorationUUID model.UUID, name string, icon string) (ObjectGroup, error) {
	guid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// language=PostgreSQL
	sql := `INSERT INTO ws_object_groups ("uuid","explorationUUID", "name", "icon", "authorUUID", "authorName") VALUES ($1,$2,$3,$4,$5,$6)
RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols),
		guid.String(), explorationUUID.String(), name, icon, model.NullUUID.String(), "")
	return ScanPgxRow(row)
}

func Edit(guid model.UUID, name string, icon string, sortWeight int) (ObjectGroup, error) {
	// language=PostgreSQL
	sql := `UPDATE ws_object_groups SET ("name", "icon", "sortWeight", "dateEdit") = ($2,$3,$4,now()) WHERE "uuid" = $1 RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String(), name, icon, sortWeight)
	return ScanPgxRow(row)
}

func Delete(guid model.UUID) (ObjectGroup, error) {
	// language=PostgreSQL
	sql := `UPDATE ws_object_groups SET "deleted" = true WHERE "uuid" = $1 RETURNING %s`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String())
	return ScanPgxRow(row)
}

func Get(guid model.UUID) (ObjectGroup, error) {
	// language=PostgreSQL
	sql := `SELECT %s FROM ws_object_groups WHERE "uuid" = $1 AND deleted = false`
	row := r.Pool.QueryRow(r.Ctx, fmt.Sprintf(sql, fullCols), guid.String())
	return ScanPgxRow(row)
}
