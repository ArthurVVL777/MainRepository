package taskService

// TaskService предоставляет методы для работы с задачами.
type TaskService struct {
	repo TaskRepository // Интерфейс для работы с репозиторием задач.
}

// NewService создает новый экземпляр сервиса задач.
func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask создает новую задачу и сохраняет ее в репозитории.
func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

// GetAllTasks возвращает все задачи из репозитория.
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

// PatchTask обновляет существующую задачу по ID.
func (s *TaskService) PatchTask(id uint, task Task) (Task, error) {
	return s.repo.PatchTaskByID(id, task)
}

// DeleteTask удаляет задачу по ID.
func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
