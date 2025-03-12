package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/seheraksam/blockchain-with-go/models"
)

func CalculateHash(block models.Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
