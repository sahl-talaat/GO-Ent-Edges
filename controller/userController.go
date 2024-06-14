package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"
	"test/ent/user"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := config.Client.User.Query().WithCard().WithGroups().WithPets().WithSpouse().WithFollowers().WithFollowing().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	usr, err := /* config.Client.User.Get(c, userID) */ config.Client.User.Query().Where(user.ID(userID)).WithCard().WithGroups().WithPets().WithSpouse().WithFollowers().WithFollowing().Only(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, usr)
}

func CreateUser(c *gin.Context) {
	var user ent.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err = config.Client.User.Create().SetName(user.Name).SetAge(user.Age).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	var user ent.User
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err = config.Client.User.UpdateOneID(userID).SetName(user.Name).SetAge(user.Age).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	err = config.Client.User.DeleteOneID(userID).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "User deleted successfully"})
}

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

func FollowUser(c *gin.Context) {
	followerID, err := strconv.Atoi(c.Param("followerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follower ID"})
		return
	}

	followeeID, err := strconv.Atoi(c.Param("followeeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followee ID"})
		return
	}

	follower, err := config.Client.User.Get(c, followerID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Follower not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	followee, err := config.Client.User.Get(c, followeeID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Followee not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	err = config.Client.User.UpdateOne(follower).
		AddFollowing(followee).
		Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User followed successfully"})
}

func UnfollowUser(c *gin.Context) {
	followerID, err := strconv.Atoi(c.Param("followerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follower ID"})
		return
	}

	followeeID, err := strconv.Atoi(c.Param("followeeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followee ID"})
		return
	}

	follower, err := config.Client.User.Get(c, followerID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Follower not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	followee, err := config.Client.User.Get(c, followeeID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Followee not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	err = config.Client.User.UpdateOne(follower).
		RemoveFollowing(followee).
		Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User unfollowed successfully"})
}
