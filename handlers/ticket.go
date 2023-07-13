package handlers

import (
	dto "landtick/dto/result"
	ticketdto "landtick/dto/ticket"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTicket struct {
	TicketRepository repositories.TicketRepository
}

func HandlerTicket(TicketRepository repositories.TicketRepository) *handlerTicket {
	return &handlerTicket{TicketRepository}
}

func (h *handlerTicket) FindTickets(c echo.Context) error {

	tickets, err := h.TicketRepository.FindTickets()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Message: "success", Data: tickets})
}

func (h *handlerTicket) GetTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ticket, err := h.TicketRepository.GetTicket(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Message: "success", Data: convertResponseTicket(ticket)})
}

func (h *handlerTicket) SearchTicket(c echo.Context) error {
	date := c.QueryParam("date")
	startStation := c.QueryParam("startStation")
	destination := c.QueryParam("destination")

	tickets, err := h.TicketRepository.SearchTicket(date, startStation, destination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Message: "success", Data: tickets})
}

func (h *handlerTicket) CreateTicket(c echo.Context) error {

	request := new(ticketdto.TicketRequest)

	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	validation := validator.New()
	error := validation.Struct(request)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	userId := c.Get("userLogin").(jwt.MapClaims)["id"].(float64)

	ticket := models.Ticket{
		NameTrain:            request.NameTrain,
		TypeTrain:            request.TypeTrain,
		StartDate:            request.StartDate,
		StartStationID:       request.StartStationID,
		StartTime:            request.StartTime,
		DestinationStationID: request.DestinationStationID,
		ArrivalTime:          request.ArrivalTime,
		Price:                request.Price,
		Qty:                  request.Qty,
		UserID:               int(userId),
	}

	ticket, err = h.TicketRepository.CreateTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	ticket, _ = h.TicketRepository.GetTicket(ticket.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Message: "success", Data: convertResponseTicket(ticket)})
}

func convertResponseTicket(u models.Ticket) ticketdto.TicketResponse {
	return ticketdto.TicketResponse{
		NameTrain:          u.NameTrain,
		TypeTrain:          u.TypeTrain,
		StartDate:          u.StartDate,
		StartStation:       u.StartStation,
		StartTime:          u.StartTime,
		DestinationStation: u.DestinationStation,
		ArrivalTime:        u.ArrivalTime,
		Price:              u.Price,
		Qty:                u.Qty,
		User:               u.User,
	}
}
