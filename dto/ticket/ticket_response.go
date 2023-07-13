package ticketdto

import "landtick/models"

type TicketResponse struct {
	NameTrain            string         `json:"name_train"`
	TypeTrain            string         `json:"type_train"`
	StartDate            string         `json:"start_date"`
	StartStationID       int            `json:"-"`
	StartStation         models.Station `json:"start_station" gorm:"foreignkey:StartStationID"`
	StartTime            string         `json:"start_time"`
	DestinationStationID int            `json:"-"`
	DestinationStation   models.Station `json:"destination_station" gorm:"foreignkey:DestinationStationID"`
	ArrivalTime          string         `json:"arival_time"`
	Price                int            `json:"price"`
	Qty                  int            `json:"qty"`
	UserID               int            `json:"-"`
	User                 models.UserResponse   `json:"user" gorm:"foreignkey:UserID"`
}