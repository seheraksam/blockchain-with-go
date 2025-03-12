package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/blockchain-with-go/initializers"
)

func init() {
	if err := initializers.LoadEnvVariables(); err != nil {
		log.Fatal("Env yüklenemedi:", err)
	}
}

func main() {

}
func run() error {
	r := gin.Default()
	httpAddr := os.Getenv("PORT")
	log.Println("Listening on ", os.Getenv("ADDR"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
