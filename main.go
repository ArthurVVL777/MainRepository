package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetMessageHandler(c echo.Context) error {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not find the tasks",
		})
	}
	var messages []string
	for _, task := range tasks {
		messages = append(messages, task.Message)
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostHandler(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not add the task",
		})
	}
	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not create the task",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Task was successfully created",
	})
}

func UpdateHandler(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid input",
		})
	}

	if err := db.Model(&Task{}).Where("id = ?", task.ID).Updates(task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "Error updating task",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Task updated successfully",
	})
}

func PatchHandler(c echo.Context) error {
	IDParam := c.Param("id")
	id, err := strconv.Atoi(IDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Bad ID",
		})
	}
	var updateTask Task
	if err := c.Bind(&updateTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid input",
		})
	}
	if err := db.Model(&Task{}).Where("id = ?", id).Updates(updateTask).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not update the task",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Task was updated",
	})
}

func DeleteHandler(c echo.Context) error {
	IDParam := c.Param("id")
	id, err := strconv.Atoi(IDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Bad ID",
		})
	}
	if err := db.Delete(&Task{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not delete the task",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Task was deleted",
	})
}

func main() {
	initDB()
	e := echo.New()
	e.GET("/tasks", GetMessageHandler)
	e.POST("/tasks", PostHandler)
	e.PUT("/tasks", UpdateHandler) // Используем PUT для обновления задачи
	e.PATCH("/tasks/:id", PatchHandler)
	e.DELETE("/tasks/:id", DeleteHandler)

	e.Start(":8080")
}
