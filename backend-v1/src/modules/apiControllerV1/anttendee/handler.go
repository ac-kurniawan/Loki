package apiControllerV1

import (
	"antriin/src/business/attendee"
	"antriin/src/util/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IAttendeeHandler interface {
	getAttendees(ctx echo.Context) error
	setAttendee(ctx echo.Context) error
 }

type Handler struct {
	attendeeService attendee.IAttendee
}

func (h Handler) setAttendee(ctx echo.Context) error {
	req := new(SetAttendeeRequest)
	if err := ctx.Bind(&req); err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusBadRequest,
			response.Error(http.StatusBadRequest, nil, &msg))
	}

	var bufferContacts []attendee.Contact
	for _,e := range req.Contacts {
		bufferContacts = append(bufferContacts, attendee.Contact{ContactName: e.ContactName, Contact: e.Contact})
	}

	attd := attendee.Attendee{
		Name: req.Name,
		EventId: req.EventId,
		Contacts: bufferContacts,
	}
	err := h.attendeeService.SetAttendee(attd)

	if err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusInternalServerError,
			response.Error(http.StatusInternalServerError, nil, &msg))
	}

	payload := NewGetAttendeesResponse([]attendee.Attendee{attd})

	return ctx.JSON(http.StatusCreated,
		response.Success(200, payload))
}

func (h Handler) getAttendees(ctx echo.Context) error {
	page := uint(1)
	err := echo.QueryParamsBinder(ctx).Uint("page", &page).BindError()

	if err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusBadRequest,
			response.Error(http.StatusBadRequest, nil, &msg))
	}

	res, err := h.attendeeService.GetAttendees(int(page))

	if err != nil {
		msg := err.Error()
		return ctx.JSON(http.StatusInternalServerError,
			response.Error(http.StatusInternalServerError, nil, &msg))
	}

	payload := NewGetAttendeesResponse(res)

	return ctx.JSON(200, response.Success(200, payload))

}

func NewAttendeeHandler(attendeeService attendee.IAttendee) IAttendeeHandler {
	return &Handler{attendeeService}
}