package blockchain

import (
	"strings"
	"errors"
	"time"
	"github.com/gunbos1031/arkhon/utils"
	"github.com/gunbos1031/arkhon/db"
)

type block struct {
	Hash			string	`json:"hash"`
	PrevHash		string	`json:"prevHash"`
	Timestamp		int		`json:"timestamp"`
	Difficulty		int		`json:"difficulty"`
	Nonce			int		`json:"nonce"`
	Height			int		`json:"height"`
	Transactions	[]*Tx	`json:"transaction"`
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

func (b *block) txsToConfirm() {
	mem := Mempool()
	txs := []*Tx{}
	coinBaseTx := makeCoinbaseTx()
	txs = append(txs, coinBaseTx)
	for _, tx := range mem.Txs {
		txs = append(txs, tx)
	}
	b.Transactions = txs
	mem.Txs = make(map[string]*Tx)
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

func createBlock(hash string, height, diff int) *block {
	b := &block{
		PrevHash: hash,
		Difficulty: diff,
		Nonce: 0,
		Height: height + 1,
	}
	b.mine()
	b.txsToConfirm()
	persistBlock(b)
	return b
}