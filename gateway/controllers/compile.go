package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sudarshan-PR/playground/gateway/models"
	"github.com/gin-gonic/gin"
)

func CompileHandler(c *gin.Context) {
	var body models.CompileBody

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Post Data"})
	}

	queueContent, err := json.Marshal(body) // Converting struct to json bytes
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while processing code"})
	}

	if err = models.PushToQueue(body.Language, "amq.direct", queueContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while pushing code to queue"})
	}

	c.JSON(http.StatusOK, gin.H{
	  "message": "Code Queued.",
	})
}
