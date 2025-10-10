package integration

import (
	"reflect"
	"testing"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
)

type WorkflowServiceTest struct {
	workflowService *services.WorkflowService
}

const pageSizeSeveralElements int = 2
const pageSizeOneElement int = 1
const secondWorkflowIndex int = 1

var expectedWorkflows = []models.Workflow{
	{Name: "first_test"},
	{Name: "second_test"},
}

func NewWorkflowServiceTest(workflowService *services.WorkflowService) *WorkflowServiceTest {
	return &WorkflowServiceTest{workflowService: workflowService}
}

// TODO: need to change according to pagination using filter
func (workflowServiceTest *WorkflowServiceTest) testCrdOperations(test *testing.T) {
	for _, workflow := range expectedWorkflows {
		workflowServiceTest.testSaveWorkflow(test, workflow)
	}
	workflowServiceTest.testGetWorkflowsByPaginationWithoutCursor(test, pageSizeSeveralElements, expectedWorkflows)
	workflowServiceTest.testGetWorkflowsByPaginationWithoutCursor(test, pageSizeOneElement, expectedWorkflows)
	gotWorkflowsByPaginationWithoutCursor, err :=
		workflowServiceTest.workflowService.GetWorkflowsByPagination("", pageSizeSeveralElements)
	if err != nil {
		test.Errorf("the issue during getting workflows without cursor")
	}
	cursorSecondElement := gotWorkflowsByPaginationWithoutCursor[secondWorkflowIndex].WorkflowId
	workflowServiceTest.testGetWorkflowsByPaginationUsingCursor(test,
		cursorSecondElement.String(), pageSizeOneElement, expectedWorkflows[secondWorkflowIndex:])
	for _, workflow := range gotWorkflowsByPaginationWithoutCursor {
		workflowServiceTest.testRemoveWorkflowById(test, workflow)
	}
}

func (workflowServiceTest *WorkflowServiceTest) testSaveWorkflow(
	test *testing.T, workflow models.Workflow) {
	err := workflowServiceTest.workflowService.SaveWorkflow(workflow)
	assertEqualNil(test, err, "the error during saving the workflow")
}

func (workflowServiceTest *WorkflowServiceTest) testGetWorkflowsByPaginationWithoutCursor(
	test *testing.T,
	pageSizeElements int, expectedWorkflows []models.Workflow) {
	gotWorkflowsByPaginationWithoutCursor, err :=
		workflowServiceTest.workflowService.GetWorkflowsByPagination("", pageSizeElements)
	assertEqualNil(test, err, "the error during getting workflows by pagination without cursor")
	assertEqual(test, len(gotWorkflowsByPaginationWithoutCursor), pageSizeElements,
		"the count of the workflows are different after getting by pagination without cursor")
	for indexWorkflow, workflow := range gotWorkflowsByPaginationWithoutCursor {
		assertEqual(test, expectedWorkflows[indexWorkflow].Name, workflow.Name,
			"the data between expected and actual workflows are different")
	}
}

func (workflowServiceTest *WorkflowServiceTest) testGetWorkflowsByPaginationUsingCursor(
	test *testing.T, cursor string, pageSizeElements int, expectedWorkflows []models.Workflow) {
	gotWorkFLowByPaginationUsingCursor, err :=
		workflowServiceTest.workflowService.GetWorkflowsByPagination(cursor, pageSizeElements)
	assertEqualNil(test, err, "the error during getting workflows by pagination using cursor")
	assertEqual(test, len(gotWorkFLowByPaginationUsingCursor), pageSizeElements,
		"the count of the workflows are different after getting by pagination using cursor")
	for indexWorkflow, workflow := range gotWorkFLowByPaginationUsingCursor {
		assertEqual(test, expectedWorkflows[indexWorkflow].Name, workflow.Name,
			"the data between expected and actual workflows are different")
	}
}

func (workflowServiceTest *WorkflowServiceTest) testRemoveWorkflowById(test *testing.T, removingWorkflow models.Workflow) {
	err := workflowServiceTest.workflowService.RemoveWorkflowById(removingWorkflow.WorkflowId.String())
	assertEqualNil(test, err, "the error during removing the workflow by id")
	_, err = workflowServiceTest.workflowService.GetWorkflowById(removingWorkflow.WorkflowId.String())
	assertEqualNotNil(test, err, "the workflow by id was not deleted")
}

func assertEqualNil(test *testing.T, err error, message string) {
	if err != nil {
		test.Errorf("%s", message)
	}
}

func assertEqualNotNil(test *testing.T, err error, message string) {
	if err == nil {
		test.Errorf("%s", message)
	}
}

func assertEqual(test *testing.T, expected any, actual any, message string) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		test.Errorf("the expected value and actual value have the different types")
	}
	if expected != actual {
		test.Errorf("%s", message)
	}
}
