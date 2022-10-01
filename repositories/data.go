package repositories

import (
	"tchtest/models"

	"gorm.io/gorm"
)

type DataImageRepository interface {
	FindDatas() ([]models.Data, error)
	GetData(ID int) (models.Data, error)
	CreateData(haveDataImage models.Data) (models.Data, error)
}

func RepositoryDataImage(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindDatas() ([]models.Data, error) {
	var datas []models.Data
	err := r.db.Preload("User").Find(&datas).Error

	return datas, err
}

func (r *repository) GetData(ID int) (models.Data, error) {
	var data models.Data
	err := r.db.Preload("User").First(&data, ID).Error

	return data, err
}

func (r *repository) CreateData(haveDataImage models.Data) (models.Data, error) {
	err := r.db.Create(&haveDataImage).Error

	return haveDataImage, err
}
