package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"
	"test/ent/group"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	grp, err := config.Client.Group.Query().WithUsers().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grp)
}

func CreateGroup(c *gin.Context) {
	var grp ent.Group
	err := c.BindJSON(&grp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JON Data"})
		return
	}
	newgrp, err := config.Client.Group.Create().SetName(grp.Name).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newgrp)
}

func GetGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}
	grp, err := config.Client.Group.Query().Where(group.ID(groupID)).WithUsers().Only(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grp)
}

func AddUserToGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	grp, err := config.Client.Group.Get(c, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usr, err := config.Client.User.Get(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = config.Client.Group.UpdateOne(grp).
		AddUserIDs(usr.ID).
		Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to group: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added to group successfully"})
}

func RemoveUserFromGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	grp, err := config.Client.Group.Get(c, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usr, err := config.Client.User.Get(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = config.Client.Group.UpdateOne(grp).
		RemoveUserIDs(usr.ID).
		Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove user from group: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed from group successfully"})
}
