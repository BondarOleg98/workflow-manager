package integration

import (
	"reflect"
	"testing"
	"workflowmanager/app/models"
	"workflowmanager/app/services"
)

// TODO: need to change according to pagination using filter
func testCrdOperations(test *testing.T) {
	expectedWorkflows := []models.Workflow{
		{Name: "first_test"},
		{Name: "second_test"},
	}
	for _, workflow := range expectedWorkflows {
		testSaveWorkflow(test, workflow)
	}
	pageSizeSeveralElements := 2
	testGetWorkflowsByPaginationWithoutCursor(test, pageSizeSeveralElements, expectedWorkflows)
	pageSizeOneElement := 1
	testGetWorkflowsByPaginationWithoutCursor(test, pageSizeOneElement, expectedWorkflows)
	gotWorkflowsByPaginationWithoutCursor, err :=
		services.GetWorkflowsByPagination("", pageSizeSeveralElements)
	if err != nil {
		test.Errorf("the issue during getting workflows without cursor")
	}
	secondWorkflowIndex := 1
	cursorSecondElement := gotWorkflowsByPaginationWithoutCursor[secondWorkflowIndex].WorkflowId
	testGetWorkflowsByPaginationUsingCursor(test,
		cursorSecondElement.String(), pageSizeOneElement, expectedWorkflows[secondWorkflowIndex:])
	for _, workflow := range gotWorkflowsByPaginationWithoutCursor {
		testRemoveWorkflowById(test, workflow)
	}
}

func testSaveWorkflow(test *testing.T, workflow models.Workflow) {
	err := services.SaveWorkflow(workflow)
	assertEqualNil(test, err, "the error during saving the workflow")
}

func testGetWorkflowsByPaginationWithoutCursor(test *testing.T,
	pageSizeElements int, expectedWorkflows []models.Workflow) {
	gotWorkflowsByPaginationWithoutCursor, err :=
		services.GetWorkflowsByPagination("", pageSizeElements)
	assertEqualNil(test, err, "the error during getting workflows by pagination without cursor")
	assertEqual(test, len(gotWorkflowsByPaginationWithoutCursor), pageSizeElements,
		"the count of the workflows are different after getting by pagination without cursor")
	for indexWorkflow, workflow := range gotWorkflowsByPaginationWithoutCursor {
		assertEqual(test, expectedWorkflows[indexWorkflow].Name, workflow.Name,
			"the data between expected and actual workflows are different")
	}
}

func testGetWorkflowsByPaginationUsingCursor(test *testing.T,
	cursor string, pageSizeElements int, expectedWorkflows []models.Workflow) {
	gotWorkFLowByPaginationUsingCursor, err :=
		services.GetWorkflowsByPagination(cursor, pageSizeElements)
	assertEqualNil(test, err, "the error during getting workflows by pagination using cursor")
	assertEqual(test, len(gotWorkFLowByPaginationUsingCursor), pageSizeElements,
		"the count of the workflows are different after getting by pagination using cursor")
	for indexWorkflow, workflow := range gotWorkFLowByPaginationUsingCursor {
		assertEqual(test, expectedWorkflows[indexWorkflow].Name, workflow.Name,
			"the data between expected and actual workflows are different")
	}
}

func testRemoveWorkflowById(test *testing.T, removingWorkflow models.Workflow) {
	err := services.RemoveWorkflowById(removingWorkflow.WorkflowId.String())
	assertEqualNil(test, err, "the error during removing the workflow by id")
	_, err = services.GetWorkflowById(removingWorkflow.WorkflowId.String())
	assertEqualNotNil(test, err, "the workflow by id was not deleted")
}

func assertEqualNil(test *testing.T, err error, message string) {
	if err != nil {
		test.Errorf(message)
	}
}

func assertEqualNotNil(test *testing.T, err error, message string) {
	if err == nil {
		test.Errorf(message)
	}
}

func assertEqual(test *testing.T, expected any, actual any, message string) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		test.Errorf("the expected value and actual value have the different types")
	}
	if expected != actual {
		test.Errorf(message)
	}
}
