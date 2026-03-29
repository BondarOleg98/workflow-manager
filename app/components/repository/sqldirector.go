package repository

import "workflowmanager/app/models/filtration"

type sqlDirector struct {
	sqlBuilder sqlBuilder
}

func newSqlDirector(sqlBuilder sqlBuilder) *sqlDirector {
	return &sqlDirector{
		sqlBuilder: sqlBuilder,
	}
}

func (sqlDirector *sqlDirector) setBuilder(sqlBuilder sqlBuilder) {
	sqlDirector.sqlBuilder = sqlBuilder
}

func (sqlDirector *sqlDirector) buildSqlRequest(tableName string,
	filtration filtration.Filtration) string {
	sqlDirector.sqlBuilder.setTableName(tableName)
	sqlDirector.sqlBuilder.setFilterFieldName(filtration.Filter.Field)
	sqlDirector.sqlBuilder.setFilterOperator(string(filtration.Filter.Operator))
	sqlDirector.sqlBuilder.setFilterValue(filtration.Filter.Value)
	sqlDirector.sqlBuilder.setSorterOperator(string(filtration.Sorter.Order))
	sqlDirector.sqlBuilder.setSorterFieldName(filtration.Sorter.Field)
	return sqlDirector.sqlBuilder.getSqlRequest()
}
