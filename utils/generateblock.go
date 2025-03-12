package utils

import (
	"time"

	"github.com/seheraksam/blockchain-with-go/models"
)

func GenerateBlock(oldBlock models.Block, BPM int) (models.Block, error) {
	var newBlock models.Block

	t := time.Now()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Timestamp = t.String()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock, nil
}
