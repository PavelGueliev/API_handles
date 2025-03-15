package handlers

import (
	"1/internal/taskService" // Импортируем наш сервис
	"1/internal/web/tasks"
	"golang.org/x/net/context"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func NewTaskHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) PatchTasks(_ context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {
	// Получаем ID из запроса
	id := request.Id

	// Декодируем тело запроса
	taskRequest := request.Body

	// Создаем структуру обновления
	updateData := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	// Обновляем задачу через сервис
	updatedTask, err := h.Service.UpdateTaskById(uint(id), updateData)
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	response := tasks.PatchTasks200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	} // Возвращаем обновленную задачу
	return response, nil
}

func (h *Handler) DeleteTasks(_ context.Context, request tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	// Получаем ID задачи
	id := request.Id

	// Вызываем сервис для удаления
	err := h.Service.DeleteTaskById(uint(id))
	if err != nil {
		return tasks.DeleteTasks404Response{}, err
	}

	// Успешное удаление - просто возвращаем пустой ответ (204 No Content)
	return tasks.DeleteTasks204Response{}, nil
}
