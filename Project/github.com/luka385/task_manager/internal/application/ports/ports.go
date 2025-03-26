package ports

import "github.com/luka385/task_manager/internal/domain"

type TaskRepository interface {
	CreateTask(*domain.Task) error
	GetTask(string) (*domain.Task, error)
	GetAllTasks() ([]*domain.Task, error)
	UpdateTask(*domain.Task) error
	DeleteTask(string) error
}
