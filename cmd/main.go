package main

import (
	"calendar-app/internal/controller"
	"calendar-app/internal/repository"
	"calendar-app/internal/service"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	dsn := "host=db user=postgres password=htmlneyazik dbname=calendar port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Инициализация репозитория, сервиса и контроллера
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	// Настройка роутера
	r := gin.Default()

	// Маршруты
	r.POST("/events", eventController.CreateEvent)
	r.PUT("/events", eventController.UpdateEvent)
	r.DELETE("/events/:id", eventController.DeleteEvent)
	r.GET("/events/:id", eventController.GetEventByID)
	r.GET("/events", eventController.GetAllEvents)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
