package queries

const InsertTaskQuery = "INSERT INTO tasks (workflow_id, name, created_at, updated_at, state) VALUES ($1, $2, $3, $4, $5)"
const RemoveTaskByIdQuery = "DELETE FROM tasks WHERE task_id = $1"
const GetTaskByIdQuery = "SELECT task_id, name, created_at, updated_at, state FROM tasks WHERE task_id = $1"
const GetTasksByWorkflowIdQuery = "SELECT task_id, name, created_at, updated_at, state FROM tasks WHERE workflow_id = $1"
const GetTasksByPaginationQuery = "SELECT task_id, name, created_at, updated_at, state FROM tasks WHERE task_id >= $1 ORDER BY task_id LIMIT $2"
const GetTasksByPaginationWithoutCursorQuery = "SELECT task_id, name, created_at, updated_at, state FROM tasks ORDER BY task_id LIMIT $1"
