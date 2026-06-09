package repository

import (
	"fmt"
)

type SelectSqlBuilder struct {
	tableName       string
	sorterFieldName string
	filterFieldName string
	filterValue     string
	filterOperator  string
	sorterOperator  string
}

func NewSelectSqlBuilder() *SelectSqlBuilder {
	return &SelectSqlBuilder{}
}

func (selectQueryBuilder *SelectSqlBuilder) setTableName(tableName string) {
	selectQueryBuilder.tableName = tableName
}

func (selectQueryBuilder *SelectSqlBuilder) setFilterOperator(filterOperator string) {
	selectQueryBuilder.filterOperator = filterOperator
}

func (selectQueryBuilder *SelectSqlBuilder) setFilterFieldName(filterFieldName string) {
	selectQueryBuilder.filterFieldName = filterFieldName
}

func (selectQueryBuilder *SelectSqlBuilder) setFilterValue(filterValue string) {
	selectQueryBuilder.filterValue = filterValue
}

func (selectQueryBuilder *SelectSqlBuilder) setSorterOperator(sorterOperator string) {
	selectQueryBuilder.sorterOperator = sorterOperator
}

func (selectQueryBuilder *SelectSqlBuilder) setSorterFieldName(sorterFieldName string) {
	selectQueryBuilder.sorterFieldName = sorterFieldName
}

func (selectQueryBuilder *SelectSqlBuilder) getSqlRequest() string {
	var prebuiltFilterOperatorWithValue string
	switch selectQueryBuilder.filterOperator {
	case "IN":
		prebuiltFilterOperatorWithValue = fmt.Sprintf("%s (%s)",
			selectQueryBuilder.filterOperator, selectQueryBuilder.filterValue)
	case "EQ":
		prebuiltFilterOperatorWithValue = fmt.Sprintf("='%s'", selectQueryBuilder.filterValue)
	case "LIKE":
		prebuiltFilterOperatorWithValue = fmt.Sprintf("%s '%s'", selectQueryBuilder.filterOperator,
			selectQueryBuilder.filterValue)
	}
	return fmt.Sprintf("SELECT * FROM %s WHERE %s::TEXT %s ORDER BY %s %s",
		selectQueryBuilder.tableName,
		selectQueryBuilder.filterFieldName,
		prebuiltFilterOperatorWithValue,
		selectQueryBuilder.sorterFieldName,
		selectQueryBuilder.sorterOperator,
	)
}
