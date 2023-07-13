package authdto

type LoginResponse struct {
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"-"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
