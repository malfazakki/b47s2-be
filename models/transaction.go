package models

type Transaction struct {
	ID       int            `json:"id" gorm:"primary_key"`
	UserID   int            `json:"user_id" form:"user_id" gorm:"type: int"`
	User     UserResponse   `json:"user"`
	TicketID int            `json:"ticket_id" form:"ticket_id" gorm:"type: int"`
	Ticket   TicketResponse `json:"ticket"`
	Status   string         `json:"status" form:"status" gorm:"type: varchar(255)"`
}
