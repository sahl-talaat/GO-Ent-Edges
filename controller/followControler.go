package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"

	"github.com/gin-gonic/gin"
)

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
