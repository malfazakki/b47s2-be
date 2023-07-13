package transactiondto

import "landtick/models"

type TransactionRequest struct {
	ID       int    `json:"id" gorm:"primary_key"`
	UserID   int    `json:"user_id" form:"user_id" gorm:"type: int" validate:"required"`
	TicketID int    `json:"ticket_id" form:"ticket_id" gorm:"type: int" validate:"required"`
	Status   string `json:"status" form:"status" gorm:"type: varchar(255)" validate:"required"`
}

type TransactionResponse struct {
	ID       int                   `json:"id"`
	UserID   int                   `json:"user_id"`
	User     models.UserResponse   `json:"user"`
	TicketID int                   `json:"ticket_id"`
	Ticket   models.TicketResponse `json:"ticket"`
	Status   string                `json:"status"`
}
