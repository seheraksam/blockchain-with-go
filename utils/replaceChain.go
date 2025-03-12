package utils

import "github.com/seheraksam/blockchain-with-go/models"

func ReplaceChain(newBlocks []models.Block) {
	if len(newBlocks) > len(models.Blockchain) {
		models.Blockchain = newBlocks
	}
}
