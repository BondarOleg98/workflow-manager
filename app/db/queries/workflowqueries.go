package queries

const GetWorkflowsQuery = "SELECT * FROM workflows"
const GetWorkflowByIdQuery = "SELECT * FROM workflows WHERE workflow_id = $1"
const RemoveWorkflowByIdQuery = "DELETE FROM workflows WHERE workflow_id = $1"
