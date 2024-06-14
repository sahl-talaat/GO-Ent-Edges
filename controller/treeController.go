package controller

import (
	"net/http"
	"strconv"
	"test/config"
	"test/ent"
	"test/ent/tree"

	"github.com/gin-gonic/gin"
)

func CreateTreeItem(c *gin.Context) {
	var input ent.Tree
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	newNode, err := config.Client.Tree.
		Create().
		SetValue(input.Value).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create node: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, newNode)
}

func GetTreeItem(c *gin.Context) {
	nodeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node ID"})
		return
	}

	node, err := config.Client.Tree.
		Query().
		Where(tree.ID(nodeID)).
		WithChildren().
		Only(c)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, node)
}

func SetParent(c *gin.Context) {
	nodeID, err := strconv.Atoi(c.Param("nodeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node ID"})
		return
	}

	parentID, err := strconv.Atoi(c.Param("parentID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node ID"})
		return
	}

	parent, err := config.Client.Tree.Get(c, parentID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	node, err := config.Client.Tree.Get(c, nodeID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	_, err = config.Client.Tree.UpdateOne(node).SetParent(parent).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set parent: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Parent set successfully"})
}

func GetLeafNodes(c *gin.Context) {

	leafNodes, err := config.Client.Tree.
		Query().
		Where(tree.Not(tree.HasChildren())).
		All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leafNodes)
}
