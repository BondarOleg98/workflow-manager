package queries

// InsertTaskQuery
// const GetTasksByPaginationQuery = "SELECT * FROM tasks WHERE task_id >= $1 ORDER BY task_id ASC LIMIT $2"
// const GetTasksByPaginationWithoutCursorQuery = "SELECT * FROM tasks ORDER BY task_id ASC LIMIT $1"
// const GetTaskByIdQuery = "SELECT * FROM tasks WHERE task_id = $1"
// const RemoveTaskByIdQuery = "DELETE FROM tasks WHERE task_id = $1"
const InsertTaskQuery = "INSERT INTO tasks (task_id, workflow_id, name, created_at, updated_at, state) VALUES ($1, $2, $3, $4, $5, $6)"
