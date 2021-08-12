package blockchain

import (
	"sync"
	"github.com/gunbos1031/arkhon/utils"
	"github.com/gunbos1031/arkhon/db"
)

type blockchain struct {
	NewestHash 			string	`json:"newestHash"`
	Height 				int		`json:"height"`
	CurrentDifficulty 	int		`json:"currentDifficulty`
}

var once sync.Once
var b *blockchain

const (
	defaultDifficulty 	int = 3
	blockInterval 		int = 2
	difficultyInterval 	int = 5
	allowedRange 		int = 2
)


func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height, getDifficulty())
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(data, b)
}

func Blocks(b *blockchain) []*block {
	var blocks []*block
	hashCursor := b.NewestHash
	for {
		b, err := FindBlock(hashCursor)
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
			b.AddBlock()
		} else {
			b.restore(checkpoint)
		}
	})
	return b
}

func getDifficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % 5 == 0 {
		return recaculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func recaculateDifficulty() int {
	allBlocks := Blocks(b)
	actualTime := (allBlocks[0].Timestamp/60) - (allBlocks[difficultyInterval-1].Timestamp/60)
	expectedInterval := blockInterval * difficultyInterval
	if actualTime < (expectedInterval - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedInterval + allowedRange) {
		return b.CurrentDifficulty - 1
	} else {
		return b.CurrentDifficulty
	}
}

func persistBlockchain(b *blockchain) {
	data := utils.ToBytes(b)
	db.SaveBlockchain(data)
}

func txs(b *blockchain) []*Tx {
	var txs []*Tx
	for _, block := range Blocks(b) {
		txs = append(txs, block.Transactions...)
	}
	return txs
}

func findTx(b *blockchain, targetId string) *Tx {
	for _, tx := range txs(b) {
		if tx.Id == targetId {
			return tx
		}
	}
	return nil
}