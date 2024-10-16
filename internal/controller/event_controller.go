package controller

import (
	"calendar-app/internal/service"
	"calendar-app/pkg/dto"
	"calendar-app/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) *EventController {
	return &EventController{eventService: eventService}
}

// CreateEvent - обработчик для создания события
func (c *EventController) CreateEvent(ctx *gin.Context) {
	var eventDTO dto.EventDTO
	if err := ctx.ShouldBindJSON(&eventDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Маппинг DTO в доменную сущность
	event := model.Event{
		Name:            eventDTO.Name,
		Place:           eventDTO.Place,
		Comment:         eventDTO.Comment,
		RecipientEmails: eventDTO.RecipientEmails,
		ApplicantEmail:  eventDTO.ApplicantEmail,
		StartDate:       eventDTO.StartDate,
		EndDate:         eventDTO.EndDate,
		IsFullDay:       eventDTO.IsFullDay,
		IsOnline:        eventDTO.IsOnline,
		AuthorEmail:     eventDTO.AuthorEmail,
	}

	if err := c.eventService.CreateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Маппинг доменной сущности в DTO для ответа
	responseDTO := dto.EventResponseDTO{
		ID:              event.ID,
		Name:            eventDTO.Name,
		Place:           eventDTO.Place,
		Comment:         eventDTO.Comment,
		RecipientEmails: eventDTO.RecipientEmails,
		ApplicantEmail:  eventDTO.ApplicantEmail,
		StartDate:       eventDTO.StartDate,
		EndDate:         eventDTO.EndDate,
		IsFullDay:       eventDTO.IsFullDay,
		IsOnline:        eventDTO.IsOnline,
		AuthorEmail:     eventDTO.AuthorEmail,
		CreatedAt:       event.CreatedAt,
		UpdatedAt:       event.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, responseDTO)
}

// UpdateEvent - обработчик для обновления события
func (c *EventController) UpdateEvent(ctx *gin.Context) {
	var eventDTO dto.EventDTO
	if err := ctx.ShouldBindJSON(&eventDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Маппинг DTO в доменную сущность
	event := model.Event{
		ID:              eventDTO.ID,
		Name:            eventDTO.Name,
		Place:           eventDTO.Place,
		Comment:         eventDTO.Comment,
		RecipientEmails: eventDTO.RecipientEmails,
		ApplicantEmail:  eventDTO.ApplicantEmail,
		StartDate:       eventDTO.StartDate,
		EndDate:         eventDTO.EndDate,
		IsFullDay:       eventDTO.IsFullDay,
		IsOnline:        eventDTO.IsOnline,
		AuthorEmail:     eventDTO.AuthorEmail,
	}

	if err := c.eventService.UpdateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Маппинг доменной сущности в DTO для ответа
	responseDTO := dto.EventResponseDTO{
		ID:              event.ID,
		Name:            eventDTO.Name,
		Place:           eventDTO.Place,
		Comment:         eventDTO.Comment,
		RecipientEmails: eventDTO.RecipientEmails,
		ApplicantEmail:  eventDTO.ApplicantEmail,
		StartDate:       eventDTO.StartDate,
		EndDate:         eventDTO.EndDate,
		IsFullDay:       eventDTO.IsFullDay,
		IsOnline:        eventDTO.IsOnline,
		AuthorEmail:     eventDTO.AuthorEmail,
		CreatedAt:       event.CreatedAt,
		UpdatedAt:       event.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, responseDTO)
}

// DeleteEvent - обработчик для удаления события
func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := c.eventService.DeleteEvent(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

// GetEventByID - обработчик для получения события по ID
func (c *EventController) GetEventByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	event, err := c.eventService.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Маппинг доменной сущности в DTO для ответа
	responseDTO := dto.EventResponseDTO{
		ID:              event.ID,
		Name:            event.Name,
		Place:           event.Place,
		Comment:         event.Comment,
		RecipientEmails: event.RecipientEmails,
		ApplicantEmail:  event.ApplicantEmail,
		StartDate:       event.StartDate,
		EndDate:         event.EndDate,
		IsFullDay:       event.IsFullDay,
		IsOnline:        event.IsOnline,
		AuthorEmail:     event.AuthorEmail,
		CreatedAt:       event.CreatedAt,
		UpdatedAt:       event.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, responseDTO)
}

// GetAllEvents - обработчик для получения всех событий
func (c *EventController) GetAllEvents(ctx *gin.Context) {
	events, err := c.eventService.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Маппинг всех событий в список DTO для ответа
	var responseDTOs []dto.EventResponseDTO
	for _, event := range events {
		responseDTO := dto.EventResponseDTO{
			ID:              event.ID,
			Name:            event.Name,
			Place:           event.Place,
			Comment:         event.Comment,
			RecipientEmails: event.RecipientEmails,
			ApplicantEmail:  event.ApplicantEmail,
			StartDate:       event.StartDate,
			EndDate:         event.EndDate,
			IsFullDay:       event.IsFullDay,
			IsOnline:        event.IsOnline,
			AuthorEmail:     event.AuthorEmail,
			CreatedAt:       event.CreatedAt,
			UpdatedAt:       event.UpdatedAt,
		}
		responseDTOs = append(responseDTOs, responseDTO)
	}

	ctx.JSON(http.StatusOK, responseDTOs)
}
