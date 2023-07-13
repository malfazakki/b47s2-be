package handlers

import (
	dto "landtick/dto/result"
	stationsdto "landtick/dto/station"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerStation struct {
	StationRepository repositories.StationRepository
}

type Stations struct {
	Stations interface{}
}

func HandlerStation(StationRepository repositories.StationRepository) *handlerStation {
	return &handlerStation{StationRepository}
}

func (h *handlerStation) FindStations(c echo.Context) error {
	stations, err := h.StationRepository.FindStations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Message: "success", Data: Stations{Stations: stations}})
}

func (h *handlerStation) GetStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Message: "success", Data: Stations{Stations: station}})
}

func (h *handlerStation) CreateStation(c echo.Context) error {
	request := new(stationsdto.CreateStationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	station := models.Station{
		Name: request.Name,
	}

	data, err := h.StationRepository.CreateStation(station)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Message: "success", Data: data})
}
