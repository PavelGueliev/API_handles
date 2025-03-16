package userService

import (
	"1/internal/taskService"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	DeleteUserById(id uint) error
	UpdateUserById(id uint, updatedUser User) (User, error)
	GetTasksForUser(userID uint) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) CreateUser(user User) (User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUserById(id uint) error {
	var user User

	if err := ur.db.Model(&user).Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) UpdateUserById(id uint, updatedUser User) (User, error) {
	var user User
	err := ur.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}
	err = ur.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	if err := ur.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
