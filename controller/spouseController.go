package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"
	"test/ent/user"

	"github.com/gin-gonic/gin"
)

func SetSpouse(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	spouseID, err := strconv.Atoi(c.Param("spouseID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid spouse ID"})
		return
	}
	usr, err := config.Client.User.Query().Where(user.ID(userID)).WithSpouse().Only(c)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	spose, err := config.Client.User.Query().Where(user.ID(spouseID)).WithSpouse().Only(c)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Spouse not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if usr.Edges.Spouse != nil {
		_, err = config.Client.User.UpdateOne(usr).ClearSpouse().Save(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear user's existing spouse: " + err.Error()})
			return
		}
	}

	if spose.Edges.Spouse != nil {
		_, err = config.Client.User.UpdateOne(spose).ClearSpouse().Save(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear spouse's existing spouse: " + err.Error()})
			return
		}
	}

	usr, err = config.Client.User.UpdateOne(usr).SetSpouse(spose).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set spouse for user: " + err.Error()})
		return
	}

	spose, err = config.Client.User.UpdateOne(spose).SetSpouse(usr).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set spouse for spouse: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Spouse set successfully", "user": usr, "spouse": spose})

}
