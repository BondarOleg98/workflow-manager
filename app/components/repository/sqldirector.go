package repository

import "workflowmanager/app/models/filtration"

type SqlDirector struct {
	sqlBuilder sqlBuilder
}

func NewSqlDirector(sqlBuilder sqlBuilder) *SqlDirector {
	return &SqlDirector{
		sqlBuilder: sqlBuilder,
	}
}

func (sqlDirector *SqlDirector) BuildSqlRequest(tableName string,
	filtration filtration.Filtration) string {
	sqlDirector.sqlBuilder.setTableName(tableName)
	sqlDirector.sqlBuilder.setFilterFieldName(filtration.Filter.Field)
	sqlDirector.sqlBuilder.setFilterOperator(string(filtration.Filter.Operator))
	sqlDirector.sqlBuilder.setFilterValue(filtration.Filter.Value)
	sqlDirector.sqlBuilder.setSorterOperator(string(filtration.Sorter.Order))
	sqlDirector.sqlBuilder.setSorterFieldName(filtration.Sorter.Field)
	return sqlDirector.sqlBuilder.getSqlRequest()
}
