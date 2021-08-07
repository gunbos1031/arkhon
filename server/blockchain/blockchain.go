package blockchain

import (
	"sync"
	"github.com/gunbos1031/arkhon/utils"
	"github.com/gunbos1031/arkhon/db"
)

type blockchain struct {
	NewestHash 	string	`json:"newestHash"`
	Height 		int		`json:"height"`
}

var once sync.Once
var b *blockchain

const (
	defaultDifficulty int = 2
)


func (b *blockchain) AddBlock(payload string) {
	block := createBlock(b.NewestHash, payload, b.Height, defaultDifficulty)
	b.NewestHash = block.Hash
	b.Height = block.Height
	persistBlockchain(b)
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(data, b)
}

func Blocks(b *blockchain) []*block {
	var blocks []*block
	hashCursor := b.NewestHash
	for {
		b, err := findBlock(hashCursor)
		utils.HandleErr(err)
		blocks = append(blocks, b)
		if b.PrevHash != "" {
			hashCursor = b.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{
			Height: 0,
		}
		checkpoint := db.LoadBlockchain()
		if checkpoint == nil {
			b.AddBlock("Genesis")
		} else {
			b.restore(checkpoint)
		}
	})
	return b
}

func persistBlockchain(b *blockchain) {
	data := utils.ToBytes(b)
	db.SaveBlockchain(data)
}