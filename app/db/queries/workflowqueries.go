package queries

const GetWorkflowsByPaginationQuery = "SELECT * FROM workflows WHERE workflow_id >= $1 ORDER BY workflow_id ASC LIMIT $2"
const GetWorkflowsByPaginationWithoutCursorQuery = "SELECT * FROM workflows ORDER BY workflow_id ASC LIMIT $1"
const GetWorkflowByIdQuery = "SELECT * FROM workflows WHERE workflow_id = $1"
const RemoveWorkflowByIdQuery = "DELETE FROM workflows WHERE workflow_id = $1"
const InsertWorkflowQuery = "INSERT INTO workflows (workflow_id, name, created_at, updated_at, state) VALUES ($1, $2, $3, $4, $5)"
