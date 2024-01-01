package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// swagger:parameters lead
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetDB() *gorm.DB {
	return db
}

func SetDB(database *gorm.DB) {
	db = database
}

func GetAllLeads() ([]Lead, error) {
	var leads []Lead
	if err := db.Find(&leads).Error; err != nil {
		return nil, err
	}
	return leads, nil
}

func GetLeadById(Id int64) (*Lead, error) {
	var getLead Lead
	if err := db.Where("ID=?", Id).First(&getLead).Error; err != nil {
		return nil, err
	}
	return &getLead, nil
}

func (l *Lead) CreateLead() (*Lead, error) {
	if err := db.Create(l).Error; err != nil {
		return nil, err
	}
	return l, nil
}

func DeleteLead(id int64) error {
	if err := db.Delete(&Lead{}, id).Error; err != nil {
		return err
	}
	return nil
}
