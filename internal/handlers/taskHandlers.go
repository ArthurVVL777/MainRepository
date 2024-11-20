package handlers

import (
	"github.com/labstack/echo/v4"             // Импортируем библиотеку Echo для работы с HTTP-запросами.
	"net/http"                                // Импортируем пакет для работы с HTTP-статусами.
	"pet_project_1_etap/internal/taskService" // Импортируем наш сервис задач
)

// Handler определяет структуру обработчика, который содержит сервис задач.
type Handler struct {
	Service *taskService.TaskService // Поле Service хранит указатель на экземпляр TaskService.
}

// NewHandler создает новый обработчик с заданным сервисом.
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service, // Инициализируем поле Service переданным объектом TaskService.
	}
}

// GetTasksHandler обрабатывает GET-запросы для получения всех задач.
func (h *Handler) GetTasksHandler(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks() // Вызываем метод GetAllTasks сервиса для получения всех задач.
	if err != nil {                       // Проверяем, произошла ли ошибка при получении задач.
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}) // Возвращаем статус 500 и сообщение об ошибке в формате JSON.
	}
	return c.JSON(http.StatusOK, tasks) // Возвращаем статус 200 и список задач в формате JSON.
}

// PostTaskHandler обрабатывает POST-запросы для создания новой задачи.
func (h *Handler) PostTaskHandler(c echo.Context) error {
	var task taskService.Task // Объявляем переменную для хранения новой задачи.

	if err := c.Bind(&task); err != nil { // Привязываем данные из запроса к объекту task.
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"}) // Если возникла ошибка, возвращаем статус 400 и сообщение об ошибке.
	}

	createdTask, err := h.Service.CreateTask(task) // Вызываем метод CreateTask сервиса для создания новой задачи.
	if err != nil {                                // Проверяем, произошла ли ошибка при создании задачи.
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}) // Возвращаем статус 500 и сообщение об ошибке в формате JSON.
	}

	return c.JSON(http.StatusCreated, createdTask) // Возвращаем статус 201 и созданную задачу в формате JSON.
}
