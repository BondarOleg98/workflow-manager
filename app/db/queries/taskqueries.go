package queries

const InsertTaskQuery = "INSERT INTO tasks (task_id, workflow_id, name, created_at, updated_at, state) VALUES ($1, $2, $3, $4, $5, $6)"
const RemoveTaskByIdQuery = "DELETE FROM tasks WHERE task_id = $1"
