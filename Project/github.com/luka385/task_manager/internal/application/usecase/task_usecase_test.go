package usecase_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/luka385/task_manager/internal/application/usecase"
	"github.com/luka385/task_manager/internal/domain"
	mock "github.com/luka385/task_manager/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTaskUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockTaskRepository(ctrl)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	now := time.Now()
	futureDate := now.Add(24 * time.Hour)
	pastDate := now.Add(-24 * time.Hour)

	tests := []struct {
		name     string
		setup    func()
		testFunc func() error
		wantErr  bool
	}{
		{
			name: "Crear Tarea Exitosamente",
			setup: func() {
				mockRepo.EXPECT().CreateTask(gomock.Any()).Return(nil)
			},
			testFunc: func() error {
				task := domain.Task{
					ID:          "1",
					Title:       "Aprender AWS",
					Description: "Estudiar la funcion lambda",
					Date:        futureDate,
					Completed:   false,
				}
				return taskUsecase.CreateTask(&task)
			},
			wantErr: false,
		},
		{
			name:  "Error al crear sin titulo",
			setup: func() {},
			testFunc: func() error {
				task := domain.Task{
					ID:   "2",
					Date: futureDate,
				}
				return taskUsecase.CreateTask(&task)
			},
			wantErr: true,
		},
		{
			name:  "Error al crear una tarea con fecha pasada",
			setup: func() {},
			testFunc: func() error {
				task := domain.Task{
					ID:    "3",
					Title: "task fail",
					Date:  pastDate,
				}
				return taskUsecase.CreateTask(&task)
			},
			wantErr: true,
		},
		{
			name: "Obtener una tarea existente",
			setup: func() {
				task := &domain.Task{ID: "4", Title: "testear esta app"}
				mockRepo.EXPECT().GetTask("4").Return(task, nil)
			},
			testFunc: func() error {
				_, err := taskUsecase.GetTask("4")
				return err
			},
			wantErr: false,
		},
		{
			name: "Error al obtener una tarea inexistente",
			setup: func() {
				mockRepo.EXPECT().GetTask("20").Return(&domain.Task{}, errors.New("task not found"))
			},
			testFunc: func() error {
				_, err := taskUsecase.GetTask("20")
				return err
			},
			wantErr: true,
		},
		{
			name: "Actualizar tarea exitosamente",
			setup: func() {
				task := &domain.Task{ID: "5", Title: "tarea original"}
				updatedTask := &domain.Task{ID: task.ID, Title: "tarea actualizada"}

				mockRepo.EXPECT().GetTask("5").Return(task, nil)
				mockRepo.EXPECT().UpdateTask(gomock.Eq(updatedTask)).Return(nil)
			},
			testFunc: func() error {
				updatedTask := domain.Task{ID: "5", Title: "tarea actualizada"}
				return taskUsecase.UpdateTask(&updatedTask)
			},
			wantErr: false,
		},
		{
			name: "Error al actualizar tarea",
			setup: func() {
				mockRepo.EXPECT().GetTask("30").Return(&domain.Task{}, errors.New("task not found"))
			},
			testFunc: func() error {
				task := domain.Task{ID: "30", Title: "non-existent"}
				return taskUsecase.UpdateTask(&task)
			},
			wantErr: true,
		},
		{
			name: "Completar tarea exitosamente",
			setup: func() {
				task := &domain.Task{ID: "6", Completed: false}
				mockRepo.EXPECT().GetTask("6").Return(task, nil)
				mockRepo.EXPECT().UpdateTask(gomock.Any()).Return(nil)
			},
			testFunc: func() error {
				return taskUsecase.CompleteTask("6")
			},
			wantErr: false,
		},
		{
			name: "error al completar una tarea inexistente",
			setup: func() {
				mockRepo.EXPECT().GetTask("60").Return(&domain.Task{}, errors.New("task not found"))
			},
			testFunc: func() error {
				return taskUsecase.CompleteTask("60")
			},
			wantErr: true,
		},
		{
			name: "eliminar tarea exitosamente",
			setup: func() {
				mockRepo.EXPECT().GetTask("7").Return(&domain.Task{ID: "7"}, nil)
				mockRepo.EXPECT().DeleteTask("7").Return(nil)
			},
			testFunc: func() error {
				return taskUsecase.DeleteTask("7")
			},
			wantErr: false,
		},
		{
			name: "error al eliminar una tarea inexistente",
			setup: func() {
				mockRepo.EXPECT().GetTask("90").Return(&domain.Task{}, errors.New("task not found"))
			},
			testFunc: func() error {
				return taskUsecase.DeleteTask("90")
			},
			wantErr: true,
		},
		{
			name: "obtener todas las tareas exitosamente",
			setup: func() {
				tasks := []*domain.Task{
					{ID: "1", Title: "Aprender AWS", Description: "Estudiar la funcion lambda", Date: futureDate, Completed: false},
					{ID: "2", Title: "Testear esta app", Description: "hacer todos los test que pueda haber", Date: futureDate, Completed: false},
				}
				mockRepo.EXPECT().GetAllTasks().Return(tasks, nil)
			},
			testFunc: func() error {
				tasks, err := taskUsecase.GetAllTasks()
				if err != nil {
					return err
				}
				if len(tasks) != 2 {
					return errors.New("solo se esperaban 2 tareas")
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "Error al obtener todas las tareas",
			setup: func() {
				mockRepo.EXPECT().GetAllTasks().Return(nil, errors.New("tasks not found"))
			},
			testFunc: func() error {
				tasks, err := taskUsecase.GetAllTasks()
				if err == nil {
					return errors.New("se esperaba un error, pero no se obtuvo")
				}
				if err.Error() != "tasks not found" {
					return fmt.Errorf("error inesperado: %v", err)
				}
				if tasks != nil {
					return errors.New("se esperaba un nil como respuesta")
				}
				return nil
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()           //configuramos el comportamiento de los mocks
			err := tc.testFunc() //ejecutamos la prueba
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
