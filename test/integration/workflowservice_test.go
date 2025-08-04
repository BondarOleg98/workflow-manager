package integration

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"workflowmanager/app/db"
	"workflowmanager/app/models"
	"workflowmanager/app/services"
)

// TODO: need to change according to pagination using filter
func TestCrdOperations(test *testing.T) {
	db.InitDatabaseInstance(db.Pool{
		DriverName:    "postgres",
		ConnectionUrl: "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable",
	})
	expectedFirstWorkflow := models.Workflow{
		Name: "first_test",
	}
	expectedSecondWorkflow := models.Workflow{
		Name: "second_test",
	}
	err := services.SaveWorkflow(expectedFirstWorkflow)
	assert.Nil(test, err, "the error during saving the first workflow")
	err = services.SaveWorkflow(expectedSecondWorkflow)
	assert.Nil(test, err, "the error during saving the second workflow")

	pageSizeSeveralElements := 2
	gotSeveralWorkflowsByPaginationWithoutCursor, err :=
		services.GetWorkflowsByPagination("", pageSizeSeveralElements)
	assert.Nil(test, err, "the error during getting two workflows by pagination")
	assert.Equal(test, len(gotSeveralWorkflowsByPaginationWithoutCursor), pageSizeSeveralElements)

	pageSizeElement := 1
	gotOneWorkflowByPaginationWithoutCursor, err :=
		services.GetWorkflowsByPagination("", pageSizeElement)
	assert.Nil(test, err, "the error during getting one workflow by pagination")
	assert.Equal(test, len(gotOneWorkflowByPaginationWithoutCursor), pageSizeElement)

	var gotFirstWorkflow models.Workflow
	for _, workflow := range gotOneWorkflowByPaginationWithoutCursor {
		gotFirstWorkflow, err = services.GetWorkflowById(workflow.WorkflowId.String())
		assert.Nil(test, err, "the error during getting one workflow by id")
	}

	gotSeveralWorkflowsByPaginationWithoutCursor, err =
		services.GetWorkflowsByPagination(gotFirstWorkflow.WorkflowId.String(), pageSizeSeveralElements)
	assert.Nil(test, err, "the error during getting two workflows by pagination")
	assert.Equal(test, len(gotSeveralWorkflowsByPaginationWithoutCursor), pageSizeSeveralElements)

	for _, workflow := range gotSeveralWorkflowsByPaginationWithoutCursor {
		err = services.RemoveWorkflowById(workflow.WorkflowId.String())
		assert.Nil(test, err, "the error during removing the workflow by id")
		_, err = services.GetWorkflowById(workflow.WorkflowId.String())
		assert.NotNil(test, err, "the first workflow by id was not deleted")
	}
}
