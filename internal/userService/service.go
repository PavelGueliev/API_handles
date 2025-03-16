package userService

import "1/internal/taskService"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService { return &UserService{repo: repo} }
func (us *UserService) CreateUser(user User) (User, error) {
	return us.repo.CreateUser(user)
}
func (us *UserService) UpdateUser(id uint, user User) (User, error) {
	return us.repo.UpdateUserById(id, user)
}
func (us *UserService) DeleteUser(id uint) error {
	return us.repo.DeleteUserById(id)
}
func (us *UserService) GetAllUsers() ([]User, error) {
	return us.repo.GetAllUsers()
}
func (us *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	return us.repo.GetTasksForUser(userID)
}
