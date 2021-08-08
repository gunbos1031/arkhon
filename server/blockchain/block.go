package blockchain

import (
	"strings"
	"errors"
	"time"
	"github.com/gunbos1031/arkhon/utils"
	"github.com/gunbos1031/arkhon/db"
)

type block struct {
	Hash		string	`json:"hash"`
	PrevHash	string	`json:"prevHash"`
	Timestamp	int		`json:"timestamp"`
	Difficulty	int		`json:"difficulty"`
	Nonce		int		`json:"nonce"`
	Height		int		`json:"height"`
	Payload		string	`json:"payload"`
}

var (
	ErrNotFound error = errors.New("block not found")
)

func (b *block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func FindBlock(hash string) (*block, error) {
	blockBytes := db.FindBlock(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	b := &block{}
	utils.FromBytes(blockBytes, b)
	return b, nil
}

func persistBlock(b *block) {
	k := []byte(b.Hash)
	v := utils.ToBytes(b)
	db.SaveBlock(k, v)
}

func createBlock(hash, payload string, height, diff int) *block {
	b := &block{
		Hash : "",
		PrevHash: hash,
		Difficulty: diff,
		Nonce: 0,
		Height: height + 1,
		Payload: payload,
	}
	b.mine()
	persistBlock(b)
	return b
}