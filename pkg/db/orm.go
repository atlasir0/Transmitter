package db

import (
	"gorm.io/gorm"
)


type ORM struct {
	db *gorm.DB
}


func NewORM(db *gorm.DB) *ORM {
	return &ORM{db}
}


func (o *ORM) CreateTransmission(transmission *TransmitterData) error {
	return o.db.Create(transmission).Error
}
