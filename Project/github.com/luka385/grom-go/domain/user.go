package domain

type Task struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `gorm:"not_null" json:"title"`
	Done  bool   `gorm:"not_nul" json:"done"`
}

type TaskRepository interface {
	Create(task *Task) error
	GetTaks() ([]Task, error)
	GetTaskByID(id string) (*Task, error)
	Update(id string, updatedTask *Task) error
	Delete(id string) error
}
