package repository

type TaskRepository interface {
	RemoveTaskById(taskId string) (int64, error)
}
