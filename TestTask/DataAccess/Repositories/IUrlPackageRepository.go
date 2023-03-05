package Repositories

import "TestTask/DataAccess/Entities"

type IUrlPackageRepository interface {
	GetById(id int) *Entities.UrlPackage
}
