package ws_exploration

import (
	"fmt"
)

func (filter Filter) ApplyToParams(where []string, params []interface{}) ([]string, []interface{}) {
	if filter.FilterName != "" {
		params = append(params, fmt.Sprintf("%%%s%%", filter.FilterName))
		// language=PostgreSQL prefix="SELECT * FROM ws_exploration WHERE "
		where = append(where, fmt.Sprintf(`"name" ILIKE $%d`, len(params)))
	}

	//if filter.ParentUUID != nil {
	//	params = append(params, filter.ParentUUID.String())
	//	// language=PostgreSQL prefix="SELECT * FROM files WHERE "
	//	where = append(where, fmt.Sprintf(`"ParentFile"=$%d `, len(params)))
	//}
	//
	//if filter.AuthorUUID != nil {
	//	params = append(params, filter.AuthorUUID.String())
	//	// language=PostgreSQL prefix="SELECT * FROM files WHERE "
	//	where = append(where, fmt.Sprintf(`"Author" =$%d `, len(params)))
	//}
	//
	//if filter.ExcludeStatuses != nil && len(*filter.ExcludeStatuses) > 0 {
	//	var paramMarks []string
	//	for _, state := range *filter.ExcludeStatuses {
	//		params = append(params, state)
	//		paramMarks = append(paramMarks, fmt.Sprintf(`$%d`, len(params)))
	//	}
	//	paramMarksString := strings.Join(paramMarks, ",")
	//	// language=PostgreSQL prefix="SELECT * FROM files WHERE "
	//	where = append(where, fmt.Sprintf(`"State" NOT IN (%s)`, paramMarksString))
	//}
	//
	//if filter.ExcludedUUIDs != nil && len(*filter.ExcludedUUIDs) > 0 {
	//	var placeHolders []string
	//	for _, b := range *filter.ExcludedUUIDs {
	//		params = append(params, b.String())
	//		placeHolders = append(placeHolders, fmt.Sprintf(`$%d`, len(params)))
	//	}
	//	bindString := strings.Join(placeHolders, ",")
	//	// language=PostgreSQL prefix="SELECT * FROM files WHERE "
	//	where = append(where, fmt.Sprintf(`"UUID" NOT IN (%s)`, bindString))
	//}
	return where, params
}

