package dataImagedto

type DataImageRequest struct {
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
}
