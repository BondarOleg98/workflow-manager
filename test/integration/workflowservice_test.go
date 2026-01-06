package integration

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
)

const pageSizeSeveralElements int = 2
const pageSizeOneElement int = 1

var expectedWorkflows = []models.Workflow{
	{
		Name:  "first_workflow",
		State: models.CREATED,
		Tasks: []models.Task{
			{
				Name:  "first_task",
				State: models.CREATED,
			},
			{
				Name:  "second_task",
				State: models.CREATED,
			},
		},
	},
	{
		Name:  "second_workflow",
		State: models.CREATED,
		Tasks: []models.Task{
			{
				Name:  "first_task",
				State: models.CREATED,
			},
		},
	},
}

type WorkflowServiceTest struct {
	workflowService *services.WorkflowService
}

func NewWorkflowServiceTest(workflowService *services.WorkflowService) *WorkflowServiceTest {
	return &WorkflowServiceTest{workflowService: workflowService}
}

func (workflowServiceTest *WorkflowServiceTest) testCrdOperations(test *testing.T) {
	workflowServiceTest.testSaveWorkflow(test)
	workflowServiceTest.testGetWorkflowsByPaginationWithoutCursor(test)
	workflowServiceTest.testGetWorkflowsByPaginationUsingCursor(test)
	workflowServiceTest.testRemoveWorkflowById(test)
}

func (workflowServiceTest *WorkflowServiceTest) testSaveWorkflow(test *testing.T) {
	for _, workflow := range expectedWorkflows {
		err := workflowServiceTest.workflowService.SaveWorkflow(workflow)
		assert.Nil(test, err, "the error during saving the workflow")
		gotWorkflowByPaginationWithoutCursor, err :=
			workflowServiceTest.workflowService.GetWorkflowsByPagination("", pageSizeOneElement)
		assert.Equal(test, workflow.Name, gotWorkflowByPaginationWithoutCursor[0].Name,
			"the data between expected and actual workflows are different")
		assert.Equal(test, workflow.State, gotWorkflowByPaginationWithoutCursor[0].State,
			"the data between expected and actual workflows are different")
	}
}

func (workflowServiceTest *WorkflowServiceTest) testGetWorkflowsByPaginationWithoutCursor(test *testing.T) {
	var cursor string
	gotWorkflowsByPaginationWithoutCursor, err :=
		workflowServiceTest.workflowService.GetWorkflowsByPagination(cursor, pageSizeSeveralElements)
	assert.Nil(test, err, "the error during getting workflows by pagination without cursor")
	assert.Equal(test, len(gotWorkflowsByPaginationWithoutCursor), pageSizeSeveralElements,
		"the count of the workflows are different after getting by pagination without cursor")
}

func (workflowServiceTest *WorkflowServiceTest) testGetWorkflowsByPaginationUsingCursor(test *testing.T) {
	var cursor string
	gotWorkflowsByPaginationWithoutCursor, err :=
		workflowServiceTest.workflowService.GetWorkflowsByPagination(cursor, pageSizeOneElement)
	gotWorkflowsByPaginationUsingCursor, err :=
		workflowServiceTest.workflowService.GetWorkflowsByPagination(
			gotWorkflowsByPaginationWithoutCursor[0].WorkflowId.String(), pageSizeSeveralElements)
	assert.Nil(test, err, "the error during getting workflows by pagination using cursor")
	assert.Equal(test, len(gotWorkflowsByPaginationUsingCursor), pageSizeOneElement,
		"the count of the workflows are different after getting by pagination using cursor")
	assert.Equal(test, expectedWorkflows[1].Name, gotWorkflowsByPaginationUsingCursor[0].Name)
	assert.Equal(test, expectedWorkflows[1].State, gotWorkflowsByPaginationUsingCursor[0].State)
}

func (workflowServiceTest *WorkflowServiceTest) testRemoveWorkflowById(test *testing.T) {
	var cursor string
	gotWorkflowsByPaginationUsingCursor, err :=
		workflowServiceTest.workflowService.GetWorkflowsByPagination(cursor, pageSizeSeveralElements)
	for _, workflow := range gotWorkflowsByPaginationUsingCursor {
		err = workflowServiceTest.workflowService.RemoveWorkflowById(workflow.WorkflowId.String())
		assert.Nil(test, err, "the error during removing the workflow by id")
		_, err = workflowServiceTest.workflowService.GetWorkflowById(workflow.WorkflowId.String())
		assert.NotNil(test, err, "the workflow by id was not deleted")
	}
}
