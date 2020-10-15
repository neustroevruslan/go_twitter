package models

type User struct {
	ID 			int 	`gorm:"unique" json:"id"`
	Login 		string 	`json:"login"`
	Password 	string	`json:"password"`
}

type UserValues struct {
	Login 		string
	Password 	string
}
