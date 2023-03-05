package DataFactories

import (
	"TestTask/DataAccess"
	"TestTask/DataAccess/Repositories"
	"TestTask/Domain"
	"gorm.io/gorm"
)

type DataFactory struct {
	db                *gorm.DB
	migrationService  *DataAccess.MigrationsService
	repository        Repositories.IUrlPackageRepository
	urlPackageService Domain.IUrlPackageService
	httpService       Domain.IHttpService
}

func (factory DataFactory) GetDb() (*gorm.DB, error) {
	if factory.db != nil {
		return factory.db, nil
	}

	db, err := DataAccess.NewDb()
	if err != nil {
		return nil, err
	}

	factory.db = db
	return db, nil
}

func (factory DataFactory) GetRepository() (Repositories.IUrlPackageRepository, error) {
	if factory.repository != nil {
		return factory.repository, nil
	}

	db, err := factory.GetDb()

	if err != nil {
		return nil, err
	}

	factory.repository = Repositories.NewUrlPackageRepository(db)
	return factory.repository, nil
}

func (factory DataFactory) GetMigrationsService() (*DataAccess.MigrationsService, error) {
	if factory.migrationService != nil {
		return factory.migrationService, nil
	}

	db, err := factory.GetDb()
	if err != nil {
		return nil, err
	}

	factory.migrationService = DataAccess.NewMigrationService(db)
	return factory.migrationService, nil
}

func (factory DataFactory) GetUrlPackageService() (Domain.IUrlPackageService, error) {
	if factory.urlPackageService != nil {
		return factory.urlPackageService, nil
	}

	repository, err := factory.GetRepository()
	if err != nil {
		return nil, err
	}

	factory.urlPackageService = Domain.NewUrlPackageService(repository, factory.GetHttpService())
	return factory.urlPackageService, nil
}

func (factory DataFactory) GetHttpService() Domain.IHttpService {
	if factory.httpService != nil {
		return factory.httpService
	}

	factory.httpService = Domain.NewHttpService()
	return factory.httpService
}
