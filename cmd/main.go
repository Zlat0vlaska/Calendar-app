package main

import (
	"calendar-app/internal/controller"
	"calendar-app/internal/repository"
	"calendar-app/internal/service"
	"calendar-app/pkg/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Подключение к базе данных
	dsn := "host=localhost user=postgres password=htmlneyazik dbname=calendar port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Миграции
	db.AutoMigrate(&model.Event{})

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
