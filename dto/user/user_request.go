package usersdto

type CreateUserRequest struct {
	FullName string `json:"full_name" form:"full_name" gorm:"type: varchar(255)" validate:"required"`
	Username string `json:"username" form:"username" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" form:"password" gorm:"type: varchar(255)" validate:"required"`
	Gender   string `json:"gender" form:"gender" gorm:"type: varchar(255)" validate:"required"`
	Phone    string    `json:"phone" form:"phone" gorm:"type: varchar(255)" validate:"required"`
	Address  string `json:"address" form:"address" gorm:"type: varchar(255)" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" form:"full_name" gorm:"type: varchar(255)" validate:"required"`
	Username string `json:"username" form:"username" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" form:"password" gorm:"type: varchar(255)" validate:"required"`
	Gender   string `json:"gender" form:"gender" gorm:"type: varchar(255)" validate:"required"`
	Phone    string    `json:"phone" form:"phone" gorm:"type: varchar(255)" validate:"required"`
	Address  string `json:"address" form:"address" gorm:"type: varchar(255)" validate:"required"`
}
