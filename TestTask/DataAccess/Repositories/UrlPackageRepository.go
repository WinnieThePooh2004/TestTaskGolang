package Repositories

import (
	"TestTask/DataAccess/Entities"
	"gorm.io/gorm"
)

type UrlPackageRepository struct {
	db *gorm.DB
}

func NewUrlPackageRepository(db *gorm.DB) IUrlPackageRepository {
	return UrlPackageRepository{db: db}
}

func (repository UrlPackageRepository) GetById(id int) *Entities.UrlPackage {
	var result Entities.UrlPackage
	err := repository.db.First(&result, id).Error
	if err != nil {
		return nil
	}
	return &result
}
