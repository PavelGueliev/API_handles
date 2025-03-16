package taskService

import (
	"errors"
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskById(id uint, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskById(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// CreateTask (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	if task.UserID == 0 {
		return Task{}, errors.New("user_id обязателен")
	}
	err := r.db.Create(&task).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskById(id uint, updatedTask Task) (Task, error) {
	var task Task
	err := r.db.First(&task, id)
	if err.Error != nil {
		return task, err.Error
	}
	if updatedTask.Task != "" {
		task.Task = updatedTask.Task
	}
	task.IsDone = updatedTask.IsDone

	r.db.Save(&task)
	return task, nil
}

func (r *taskRepository) DeleteTaskById(id uint) error {
	var task Task

	err := r.db.Model(&task).Where("id = ?", id).Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}
