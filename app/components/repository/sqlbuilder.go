package repository

type sqlBuilder interface {
	setTableName(tableName string)
	setFilterValue(filterName string)
	setFilterFieldName(fieldName string)
	setSorterFieldName(fieldName string)
	setFilterOperator(operator string)
	setSorterOperator(order string)
	getSqlRequest() string
}
