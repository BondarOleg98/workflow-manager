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

// TODO implement if we need to get data using EQ or IN. Also need to implement more logic for LIKE operator
func (selectQueryBuilder *selectSqlBuilder) getSqlRequest() string {
	return fmt.Sprintf("SELECT * FROM %s WHERE %s %s '%s' ORDER BY %s %s",
		selectQueryBuilder.tableName,
		selectQueryBuilder.filterFieldName,
		selectQueryBuilder.filterOperator,
		selectQueryBuilder.filterValue,
		selectQueryBuilder.sorterFieldName,
		selectQueryBuilder.sorterOperator,
	)
}
