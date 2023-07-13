package models

type Station struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type StationMyTicketResponse struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (Station) TableName() string {
	return "stations"
}
