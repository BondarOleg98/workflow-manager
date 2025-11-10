package queries

const RemoveActionsByTaskIdQuery = "DELETE FROM actions WHERE task_id IN (SELECT task_id FROM tasks WHERE workflow_id = $1)"
