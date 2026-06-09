package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/models/filtration"
)

func TestSelectBuilderWithIn(test *testing.T) {
	selectSqlBuilder := repository.NewSelectSqlBuilder()
	sqlDirector := repository.NewSqlDirector(selectSqlBuilder)
	filter := filtration.Filtration{
		Filter: filtration.Filter{
			Operator: "IN",
			Field:    "name",
			Value:    "'test_1','test_2'",
		},
		Sorter: filtration.Sorter{
			Order: "ASC",
			Field: "name",
		},
	}
	sqlExpectedQuery := "SELECT * FROM workflows WHERE name::TEXT IN ('test_1','test_2') ORDER BY name ASC"
	sqlActualQuery := sqlDirector.BuildSqlRequest("workflows", filter)
	assert.Equal(test, sqlExpectedQuery, sqlActualQuery)
}

func TestSelectBuilderWithEq(test *testing.T) {
	selectSqlBuilder := repository.NewSelectSqlBuilder()
	sqlDirector := repository.NewSqlDirector(selectSqlBuilder)
	filter := filtration.Filtration{
		Filter: filtration.Filter{
			Operator: "EQ",
			Field:    "name",
			Value:    "test",
		},
		Sorter: filtration.Sorter{
			Order: "DESC",
			Field: "name",
		},
	}
	sqlExpectedQuery := "SELECT * FROM workflows WHERE name::TEXT ='test' ORDER BY name DESC"
	sqlActualQuery := sqlDirector.BuildSqlRequest("workflows", filter)
	assert.Equal(test, sqlExpectedQuery, sqlActualQuery)
}

func TestSelectBuilderWithLike(test *testing.T) {
	selectSqlBuilder := repository.NewSelectSqlBuilder()
	sqlDirector := repository.NewSqlDirector(selectSqlBuilder)
	filter := filtration.Filtration{
		Filter: filtration.Filter{
			Operator: "LIKE",
			Field:    "name",
			Value:    "test%",
		},
		Sorter: filtration.Sorter{
			Order: "DESC",
			Field: "name",
		},
	}
	sqlExpectedQuery := "SELECT * FROM workflows WHERE name::TEXT LIKE 'test%' ORDER BY name DESC"
	sqlActualQuery := sqlDirector.BuildSqlRequest("workflows", filter)
	assert.Equal(test, sqlExpectedQuery, sqlActualQuery)
}
