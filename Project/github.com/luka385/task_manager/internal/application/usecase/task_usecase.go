package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/luka385/task_manager/internal/application/ports"
	"github.com/luka385/task_manager/internal/domain"
)

type TaskUsecase struct {
	repoTask ports.TaskRepository
}

func NewTaskUsecase(r ports.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repoTask: r}
}

func (ts *TaskUsecase) CreateTask(task *domain.Task) error {
	// verifico que el titulo no este vacio
	if task.Title == "" {
		return errors.New("the title is required")
	}
	// verifico que la fecha no sea antigua
	if task.Date.Before(time.Now()) {
		return errors.New("the deadline must be in the future")
	}
	return ts.repoTask.CreateTask(task)
}

func (ts *TaskUsecase) GetTask(id string) (*domain.Task, error) {
	task, err := ts.repoTask.GetTask(id)
	if err != nil {
		return &domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (ts *TaskUsecase) GetAllTasks() ([]*domain.Task, error) {
	tasks, err := ts.repoTask.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("error al obtener tareas: %w", err)
	}
	return tasks, nil
}

func (ts *TaskUsecase) UpdateTask(task *domain.Task) error {
	existingTask, err := ts.repoTask.GetTask(task.ID)
	if err != nil {
		return errors.New("task not found")
	}
	if task.Title == "" {
		return errors.New("the title is required")
	}
	if task.Date.IsZero() {
		if task.Date.Before(time.Now()) {
			return errors.New("the new deadline must be in the future")
		}
		existingTask.Date = task.Date
	}
	existingTask.Completed = task.Completed

	return ts.repoTask.UpdateTask(existingTask)
}

func (ts *TaskUsecase) CompleteTask(id string) error {
	task, err := ts.repoTask.GetTask(id)
	if err != nil {
		return errors.New("task not found")
	}
	task.Completed = true

	return ts.repoTask.UpdateTask(task)
}

func (ts *TaskUsecase) DeleteTask(id string) error {
	_, err := ts.repoTask.GetTask(id)
	if err != nil {
		return errors.New("task not found")
	}
	return ts.repoTask.DeleteTask(id)
}
