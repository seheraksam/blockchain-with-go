package utils

import "github.com/seheraksam/blockchain-with-go/models"

func IsBlockValid(newBlock, oldBlock models.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
