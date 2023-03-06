package DataAccess

import (
	"TestTask/DataAccess/Entities"
	"gorm.io/gorm"
	"strconv"
)

type MigrationsService struct {
	db *gorm.DB
}

func NewMigrationService(db *gorm.DB) *MigrationsService {
	return &MigrationsService{db: db}
}

func (context MigrationsService) EnsureMigrated() error {
	err := context.db.AutoMigrate(&Entities.UrlPackage{})

	if err != nil {
		return err
	}

	var count int64
	if err := context.db.Model(&Entities.UrlPackage{}).Count(&count).Error; err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	for i := 1; i < 5; i++ {
		context.db.Create(&Entities.UrlPackage{Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=" + strconv.Itoa(i)})
	}

	return nil
}
