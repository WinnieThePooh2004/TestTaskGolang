package DataAccess

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:Vova1234@localhost:1433?database=url_packages_db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
