package models

type User struct {
	ID               int64  `json:"id" gorm:"primary_key;auto_increment"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Status           string `json:"status"`
	LastConnectionIp string `json:"last_connection_ip"`
	File             string `json:"file"`
	Active           bool   `json:"active"`
	Tfa              bool   `json:"tfa"`
	LastLogin        string `json:"last_login"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
