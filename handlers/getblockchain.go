package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/blockchain-with-go/models"
)

func handleGetBlockchain(c *gin.Context) {
	bytes, err := json.MarshalIndent(models.Blockchain, "", "  ")
	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, bytes)
}
