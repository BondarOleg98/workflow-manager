package mapper

import (
	"database/sql"
	"log"
	"reflect"
)

func ListEntitiesMapped[T any](entities []T, dbRows *sql.Rows) ([]T, error) {
	for dbRows.Next() {
		var entity T
		entity, err := EntityMapped(entity, dbRows)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func EntityMapped[T any](entity T, dbRows *sql.Rows) (T, error) {
	var err error
	structValue := reflect.ValueOf(&entity).Elem()
	structType := reflect.TypeOf(entity)
	numFields := structType.NumField()
	var args []any
	for i := 0; i < numFields; i++ {
		field := structType.Field(i)
		if field.Tag.Get("mapper") != "omit" {
			args = append(args, structValue.Field(i).Addr().Interface())
		}
	}
	if err = dbRows.Scan(args...); err != nil {
		log.Fatalf("The error during mapping data from DB %s", err)
	}
	return entity, err
}
