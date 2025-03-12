package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/seheraksam/blockchain-with-go/models"
	"github.com/seheraksam/blockchain-with-go/utils"
)

func handleWriteBlock(c *gin.Context) {
	var m models.Message

	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&m); err != nil {
		c.JSON(http.StatusBadRequest, c.Body)
		return
	}
	defer c.Body.Close()

	newBlock, err := utils.GenerateBlock(models.Blockchain[len(models.Blockchain)-1], m.BPM)
	if err != nil {
		c.JSON(http.StatusInternalServerError, m)
		return
	}
	if utils.IsBlockValid(newBlock, models.Blockchain[len(models.Blockchain)-1]) {
		newBlockchain := append(models.Blockchain, newBlock)
		utils.ReplaceChain(newBlockchain)
		spew.Dump(models.Blockchain)
	}

	c.JSON(http.StatusCreated, newBlock)

}
