package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"
	"test/ent/pet"

	"github.com/gin-gonic/gin"
)

func GetPet(c *gin.Context) {
	petID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pet ID"})
		return
	}
	pt, err := config.Client.Pet.Query().Where(pet.ID(petID)).WithOwner().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pt)
}

func GetPets(c *gin.Context) {
	pets, err := config.Client.Pet.Query().WithOwner().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pets)
}

func CreatePet(c *gin.Context) {
	ownerID, err := strconv.Atoi(c.Param("ownerID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner ID"})
		return
	}
	var input ent.Pet
	err = c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body data"})
		return
	}
	pt, err := config.Client.Pet.Create().SetName(input.Name).SetAge(input.Age).SetOwnerID(ownerID).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pt)
}
