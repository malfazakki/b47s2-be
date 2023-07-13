package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	FindTickets() ([]models.Ticket, error)
	GetTicket(ID int) (models.Ticket, error)
	CreateTicket(ticket models.Ticket) (models.Ticket, error)
	SearchTicket(date string, startStation string, destinationStation string) ([]models.TicketResponse, error)
}

func RepositoryTicket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").Preload("User").Find(&tickets).Error

	return tickets, err
}

func (r *repository) GetTicket(ID int) (models.Ticket, error) {
	var ticket models.Ticket

	err := r.db.Preload("StartStation").Preload("DestinationStation").Preload("User").First(&ticket, ID).Error
	return ticket, err
}

func (r *repository) SearchTicket(date string, startStation string, destinationStation string) ([]models.TicketResponse, error) {
	var tickets []models.TicketResponse
	query := r.db.Preload("StartStation").Preload("DestinationStation")
	if date != "" {
		query = query.Where("start_date = ?", date)
	}
	if startStation != "" {
		query = query.Joins("JOIN stations AS start_station ON start_station.id = tickets.start_station_id").Where("start_station.name = ?", startStation)
	}

	if destinationStation != "" {
		query = query.Joins("JOIN stations AS destination_station ON destination_station.id = tickets.destination_station_id").Where("destination_station.name = ?", destinationStation)
	}

	err := query.Find(&tickets).Error
	if err != nil {
		return nil, err
	}

	return tickets, err
}

func (r *repository) CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	err := r.db.Create(&ticket).Error

	return ticket, err
}
