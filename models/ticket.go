package models

type Ticket struct {
	ID                   int          `json:"id" gorm:"primary_key:auto_increment"`
	NameTrain            string       `json:"name_train" form:"train_name" gorm:"type:varchar(255)"`
	TypeTrain            string       `json:"type_train" form:"type" gorm:"type:varchar(255)"`
	StartDate            string       `json:"start_date" gorm:"type:varchar(255)"`
	StartStationID       int          `json:"start_station_id" gorm:"type:int"`
	StartStation         Station      `json:"start_station" gorm:"foreignkey:StartStationID"`
	StartTime            string       `json:"start_time" gorm:"type:varchar(255)"`
	DestinationStationID int          `json:"destination_station_id" gorm:"type:int"`
	DestinationStation   Station      `json:"destination_station" gorm:"foreignkey:DestinationStationID"`
	ArrivalTime          string       `json:"arrival_time" gorm:"type:varchar(255)"`
	Price                int          `json:"price" form:"price" gorm:"type:int"`
	Qty                  int          `json:"qty" form:"qty" gorm:"type:int"`
	UserID               int          `json:"user_id" form:"user_id" gorm:"type:int"`
	User                 UserResponse `json:"user" gorm:"foreignkey:UserID"`
}

type TicketResponse struct {
	ID                   int     `json:"-"`
	NameTrain            string  `json:"name_train"`
	TypeTrain            string  `json:"type_train"`
	StartDate            string  `json:"start_date"`
	StartStationID       int     `json:"-"`
	StartStation         Station `json:"start_station"`
	StartTime            string  `json:"start_time"`
	DestinationStationID int     `json:"-"`
	DestinationStation   Station `json:"destination_station"`
	ArrivalTime          string  `json:"arrival_time"`
	Price                int     `json:"price"`
}

func (TicketResponse) TableName() string {
	return "tickets"
}
