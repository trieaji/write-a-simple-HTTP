package models

type Data struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id" form:"user_id"`
	User   User   `json:"user" form:"user"`
}

type DataImageResponse struct {
	ID     int    `json:"id"`
	Image  string `json:"image"`
	UserID int    `json:"-"`
}

func (DataImageResponse) TableName() string { //merupakan hal yang tidak perlu dimigrasi
	return "data"
}
