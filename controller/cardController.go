package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"
	"test/ent/card"

	"github.com/gin-gonic/gin"
)

func GetCard(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Card ID"})
		return
	}
	crd, err := /* config.Client.Card.Get(c, cardID) */ config.Client.Card.Query().Where(card.ID(cardID)).WithOwner().Only(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, crd)
}

func GetCards(c *gin.Context) {
	cars, err := config.Client.Card.Query().WithOwner().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, cars)
}

func CreateCard(c *gin.Context) {
	ownerID, err := strconv.Atoi(c.Param("ownerid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Owner ID"})
		return
	}
	var input ent.Card
	err = c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	card, err := config.Client.Card.Create().SetNumber(input.Number).SetPassword(input.Password).SetCach(input.Cach).SetOwnerID(ownerID).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "o2o relation " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, card)
}
