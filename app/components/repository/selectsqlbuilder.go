package repository

import (
	"fmt"
)

type selectSqlBuilder struct {
	tableName       string
	sorterFieldName string
	filterFieldName string
	filterValue     string
	filterOperator  string
	sorterOperator  string
}

func newSelectSqlBuilder() *selectSqlBuilder {
	return &selectSqlBuilder{}
}

func (selectQueryBuilder *selectSqlBuilder) setTableName(tableName string) {
	selectQueryBuilder.tableName = tableName
}

func (selectQueryBuilder *selectSqlBuilder) setFilterOperator(filterOperator string) {
	selectQueryBuilder.filterOperator = filterOperator
}

func (selectQueryBuilder *selectSqlBuilder) setFilterFieldName(filterFieldName string) {
	selectQueryBuilder.filterFieldName = filterFieldName
}

func (selectQueryBuilder *selectSqlBuilder) setFilterValue(filterValue string) {
	selectQueryBuilder.filterValue = filterValue
}

func (selectQueryBuilder *selectSqlBuilder) setSorterOperator(sorterOperator string) {
	selectQueryBuilder.sorterOperator = sorterOperator
}

func (selectQueryBuilder *selectSqlBuilder) setSorterFieldName(sorterFieldName string) {
	selectQueryBuilder.sorterFieldName = sorterFieldName
}

func (selectQueryBuilder *selectSqlBuilder) getSqlRequest() string {
	var prebuiltFilterOperatorWithValue string
	switch selectQueryBuilder.filterOperator {
	case "IN":
		prebuiltFilterOperatorWithValue = fmt.Sprintf("%s (%s)",
			selectQueryBuilder.filterOperator, selectQueryBuilder.filterValue)
	case "EQ":
		prebuiltFilterOperatorWithValue = fmt.Sprintf("='%s'", selectQueryBuilder.filterValue)
	default:
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
