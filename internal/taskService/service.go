package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}
func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}
func (s *TaskService) UpdateTaskById(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskById(id, task)
}
func (s *TaskService) DeleteTaskById(id uint) error {
	return s.repo.DeleteTaskById(id)
}
