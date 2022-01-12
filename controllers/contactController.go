package controllers

import (
	"go-contact/models"
	"go-contact/repo"
	"go-contact/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

// struct for input
type inputContact struct {
	ID     uuid.UUID `json:"id" binding:"max=63"`
	Name   string    `json:"name"`
	Gender string    `json:"gender"`
	Phone  string    `json:"phone" binding:"max=12"`
	Email  string    `json:"email"`
}

func GetAllContact(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var contacts []models.Contact

	db.Find(&contacts)

	// return json
	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

func CreateContact(c *gin.Context) {
	var input inputContact

	// Validasi inputan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid := uuid.NewV4()

	contact := models.Contact{
		ID:     uuid,
		Name:   input.Name,
		Gender: input.Gender,
		Phone:  input.Phone,
		Email:  input.Email,
	}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&contact)
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func GetContactById(c *gin.Context) {
	var contact models.Contact

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return json
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func UpdateContact(c *gin.Context) {
	// validasi inputan
	var input inputContact

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find data in database (exist or not)
	var contact models.Contact

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var updatedInputContact models.Contact

	updatedInputContact.Name = input.Name
	updatedInputContact.Gender = input.Gender
	updatedInputContact.Phone = input.Phone
	updatedInputContact.Email = input.Email
	updatedInputContact.UpdatedAt = time.Now()

	db.Model(&contact).Updates(updatedInputContact)
	// return json
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func DeleteContact(c *gin.Context) {
	var contact models.Contact

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&contact)

	// return json
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func GetAllContactPaged(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var contact models.Contact
	contactLists, err := repo.GetAllContacts(&contact, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": contactLists,
	})

}
