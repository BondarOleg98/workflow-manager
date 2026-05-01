package integration

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/components/services"
	"workflowmanager/app/db"
	"workflowmanager/app/models"
)

var expectedWorkflows = []models.Workflow{
	{
		Name:  "first_workflow",
		State: models.CREATED,
		Tasks: []models.Task{
			{Name: "first_task", State: models.CREATED},
			{Name: "second_task", State: models.CREATED},
		},
	},
	{
		Name:  "second_workflow",
		State: models.CREATED,
		Tasks: []models.Task{
			{Name: "first_task", State: models.CREATED},
		},
	},
	{
		Name:  "third_workflow",
		State: models.CREATED,
		Tasks: []models.Task{
			{Name: "third_task", State: models.CREATED},
		},
	},
}

func preconfigureTest() *services.WorkflowService {
	pool := db.Pool{
		DriverName:    "postgres",
		ConnectionUrl: "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable",
	}
	db.InitDatabaseInstance(pool)
	workflowService := initWorkflowService()
	cleanOldData(workflowService)
	return workflowService
}

func initWorkflowService() *services.WorkflowService {
	workflowRepository := repository.NewPostgresWorkflowRepository(db.GetDatabaseInstance())
	return services.NewWorkflowService(workflowRepository)
}

func cleanOldData(workflowService *services.WorkflowService) {
	workflowsForRemoving, _ := workflowService.GetWorkflowsByPagination("", len(expectedWorkflows))
	for _, workflow := range workflowsForRemoving {
		_ = workflowService.RemoveWorkflowById(workflow.WorkflowId.String())
	}
}

func createTestData(workflowService *services.WorkflowService) {
	for _, workflow := range expectedWorkflows {
		_ = workflowService.SaveWorkflow(workflow)
	}
}

func TestSaveWorkflow(test *testing.T) {
	workflowServiceTest := preconfigureTest()
	defer cleanOldData(workflowServiceTest)

	for _, workflow := range expectedWorkflows {
		err := workflowServiceTest.SaveWorkflow(workflow)
		assert.Nil(test, err, "the error during saving the workflow")
	}
	gotWorkflowByPaginationWithoutCursor, _ :=
		workflowServiceTest.GetWorkflowsByPagination("", len(expectedWorkflows))
	assert.Equal(test, expectedWorkflows[0].Name, gotWorkflowByPaginationWithoutCursor[0].Name,
		"the data between expected and actual workflows are different")
	assert.Equal(test, expectedWorkflows[0].State, gotWorkflowByPaginationWithoutCursor[0].State,
		"the data between expected and actual workflows are different")
}

func TestGetWorkflowsByPaginationWithoutCursor(test *testing.T) {
	workflowServiceTest := preconfigureTest()

	createTestData(workflowServiceTest)
	defer cleanOldData(workflowServiceTest)
	gotWorkflowsByPaginationWithoutCursor, err :=
		workflowServiceTest.GetWorkflowsByPagination("", len(expectedWorkflows))
	assert.Nil(test, err, "the error during getting workflows by pagination without cursor")
	assert.Equal(test, len(gotWorkflowsByPaginationWithoutCursor), len(expectedWorkflows),
		"the count of the workflows are different after getting by pagination without cursor")
}

func TestGetWorkflowsByPaginationUsingCursor(test *testing.T) {
	workflowServiceTest := preconfigureTest()

	const pageSizeOneElement = 1
	const pageSeveralElements = 2
	createTestData(workflowServiceTest)
	defer cleanOldData(workflowServiceTest)
	gotWorkflowsByPaginationWithoutCursor, _ :=
		workflowServiceTest.GetWorkflowsByPagination("", pageSizeOneElement)
	gotWorkflowsByPaginationUsingCursor, err :=
		workflowServiceTest.GetWorkflowsByPagination(
			gotWorkflowsByPaginationWithoutCursor[0].WorkflowId.String(), pageSeveralElements)
	assert.Nil(test, err, "the error during getting workflows by pagination using cursor")
	assert.Equal(test, len(gotWorkflowsByPaginationUsingCursor), pageSeveralElements,
		"the count of the workflows are different after getting by pagination using cursor")
	assert.Equal(test, expectedWorkflows[0].Name, gotWorkflowsByPaginationUsingCursor[0].Name)
	assert.Equal(test, expectedWorkflows[0].State, gotWorkflowsByPaginationUsingCursor[0].State)
}

func TestRemoveWorkflowById(test *testing.T) {
	workflowServiceTest := preconfigureTest()

	createTestData(workflowServiceTest)
	defer cleanOldData(workflowServiceTest)
	gotWorkflowsByPaginationUsingCursor, _ :=
		workflowServiceTest.GetWorkflowsByPagination("", len(expectedWorkflows))
	for _, workflow := range gotWorkflowsByPaginationUsingCursor {
		err := workflowServiceTest.RemoveWorkflowById(workflow.WorkflowId.String())
		assert.Nil(test, err, "the error during removing the workflow by id")
		_, err = workflowServiceTest.GetWorkflowById(workflow.WorkflowId.String())
		assert.NotNil(test, err, "the workflow by id was not deleted")
	}
}
