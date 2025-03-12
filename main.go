package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/seheraksam/blockchain-with-go/handlers"
	"github.com/seheraksam/blockchain-with-go/initializers"
	"github.com/seheraksam/blockchain-with-go/models"
)

func init() {
	if err := initializers.LoadEnvVariables(); err != nil {
		log.Fatal("Env yüklenemedi:", err)
	}
}

func main() {
	go func() {
		t := time.Now()
		genesisBlock := models.Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		models.Blockchain = append(models.Blockchain, genesisBlock)
	}()
	r := gin.Default()
	r.GET("/gethandle", handlers.HandleGetBlockchain)
	r.POST("/signUp", handlers.HandleWriteBlock)
	r.Run()

}
