package mapper

import (
	"database/sql"
	"log"
	"reflect"
	"workflowmanager/app/models"
)

func ListMapped[T any](entities []T, dbRows *sql.Rows) ([]T, error) {
	for dbRows.Next() {
		var entity T
		structValue := reflect.ValueOf(&entity).Elem()
		numFields := structValue.NumField()
		var args []interface{}
		for i := 0; i < numFields; i++ {
			field := structValue.Field(i).Interface()
			if reflect.TypeOf(field).Kind() != reflect.Slice {
				args = append(args, structValue.Field(i).Addr().Interface())
			}
		}
		if err := dbRows.Scan(args...); err != nil {
			log.Fatalf("The error during mapping data from DB %s", err)
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func WorkflowMapped(row *sql.Rows) (models.Workflow, error) {
	var err error
	workflow := models.Workflow{}
	if err = row.Scan(&workflow.WorkflowId, &workflow.Name, &workflow.CreatedAt, &workflow.UpdatedAt); err != nil {
		log.Printf("The error during mapping data from DB %s", err)
		return workflow, err
	}
	return workflow, err
}
