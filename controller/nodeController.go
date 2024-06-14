package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"
	"test/ent/node"

	"github.com/gin-gonic/gin"
)

func GetNode(c *gin.Context) {
	nodeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Node ID"})
	}
	nde, err := config.Client.Node.Query().Where(node.ID(nodeID)).WithNext().WithPrev().Only(c) /* Get(c, nodeID) */
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nde)
}

func GetNodes(c *gin.Context) {
	nodes, err := config.Client.Node.Query().WithPrev().WithNext().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func CreateNode(c *gin.Context) {
	var input ent.Node
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	lastNode, err := config.Client.Node.Query().Order(ent.Desc(node.FieldID)).First(c)
	if err != nil && !ent.IsNotFound(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get last node " + err.Error()})
		return
	}
	newNode, err := config.Client.Node.Create().SetValue(input.Value).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creating head " + err.Error()})
		return
	}
	if lastNode != nil {
		_, err = config.Client.Node.UpdateOne(lastNode).SetNext(newNode).Save(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "set new as next" + err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, newNode)
}
