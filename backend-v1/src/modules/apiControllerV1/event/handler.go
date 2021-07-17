package apiControllerV1

import (
	"antriin/src/business/event"
	"antriin/src/util/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IEventHandler interface {
	SetEvent(ctx echo.Context) error
	GetEventsByCreatorID(ctx echo.Context) error

	CreateScheduleByEvent(ctx echo.Context) error
}

type Handler struct {
	eventService event.IEvent
}

func (h Handler) CreateScheduleByEvent(ctx echo.Context) error {
	eventId := ctx.Param("eventId")
	req := new([]ScheduleRequest)
	if err := ctx.Bind(&req); err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusBadRequest,
			response.Error(http.StatusBadRequest, nil, &msg))
	}

	schedule := RequestToSchedule(*req)

	res, err := h.eventService.CreateSchedules(eventId, schedule)
	if err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusInternalServerError,
			response.Error(http.StatusInternalServerError, nil, &msg))
	}

	payload := NewEventResponse(res)
	return ctx.JSON(http.StatusOK,
		response.Success(200, payload))
}

func (h Handler) GetEventsByCreatorID(ctx echo.Context) error {
	creatorId := ctx.QueryParam("creatorId")
	result, err := h.eventService.GetEventsByCreatorID(creatorId)
	if err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusInternalServerError,
			response.Error(http.StatusInternalServerError, nil, &msg))
	}

	var payload []EventResponse
	for _, e := range result {
		payload = append(payload, NewEventResponse(e))
	}

	return ctx.JSON(http.StatusOK,
		response.Success(200, payload))
}

func (h Handler) SetEvent(ctx echo.Context) error {
	req := new(EventRequest)
	if err := ctx.Bind(&req); err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusBadRequest,
			response.Error(http.StatusBadRequest, nil, &msg))
	}

	eventData := RequestToEvent(req)
	// get userId from jwt
	userId := ctx.Get("userId").(string)

	events, err := h.eventService.SetEvent(eventData, userId)
	if err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusInternalServerError,
			response.Error(http.StatusInternalServerError, nil, &msg))
	}

	payload := NewEventResponse(events)
	return ctx.JSON(http.StatusCreated,
		response.Success(200, payload))
}

func NewEventHandler(eventService event.IEvent) IEventHandler {
	return &Handler{eventService}
}