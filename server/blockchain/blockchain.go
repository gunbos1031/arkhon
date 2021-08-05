package blockchain

import (
	"sync"
	"github.com/gunbos1031/arkhon/utils"
	"github.com/gunbos1031/arkhon/db"
)

const (
	defaultDifficulty int = 3
)

type blockchain struct {
	NewestHash 	string	`json:"newestHash"`
	Height 		int		`json:"height"`
}

var once sync.Once
var b *blockchain

func Blockchain() *blockchain {
	once.Do(func() {
		if b == nil {
			b = &blockchain{
				Height: 0,
			}
		}
	})
	return b
}

func AddBlock(payload string) {
	block := createBlock(payload, defaultDifficulty)
	blockchain := Blockchain()
	blockchain.NewestHash = block.Hash
	blockchain.Height = block.Height
	persistBlockchain(blockchain)
}

func Blocks(b *blockchain) []*block {
	var blocks []*block
	hashCursor := b.NewestHash
	for {
		b, _ := findBlock(hashCursor)
		blocks = append(blocks, b)
		if b.PrevHash != "" {
			hashCursor = b.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func persistBlockchain(b *blockchain) {
	data := utils.ToBytes(b)
	db.SaveBlockchain(data)
}