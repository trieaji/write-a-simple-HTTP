package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" form:"password" gorm:"type: varchar(255)"`
}
