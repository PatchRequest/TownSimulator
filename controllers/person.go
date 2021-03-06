package controllers

import (
	"net/http"

	"TownSimulator/models"

	"github.com/gin-gonic/gin"
)

func AllPersons(c *gin.Context) {
	var persons []models.Person
	// get all persons with all their relations
	models.DB.Preload("Jobs").Find(&persons)

	c.JSON(http.StatusOK, gin.H{"data": persons})
}

func CreatePerson(c *gin.Context) {
	var input models.CreatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person := models.Person{Name: input.Name, Surname: input.Surname, Race: input.Race, Age: input.Age, Alive: input.Alive}
	models.DB.Preload("Jobs").Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func FindPerson(c *gin.Context) {
	var person models.Person

	if err := models.DB.Preload("Jobs").Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	if err := models.DB.Preload("Jobs").Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.UpdatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&person).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": person})
}
