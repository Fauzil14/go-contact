package repo

import (
	"go-contact/config"
	"go-contact/models"
)

func GetAllContacts(user *models.Contact, pagination *models.Pagination) (*[]models.Contact, error) {
	var contacts []models.Contact
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Contact{}).Where(user).Find(&contacts)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &contacts, nil
}
