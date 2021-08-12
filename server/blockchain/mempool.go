package blockchain

import (
	"sync"
	"github.com/gunbos1031/arkhon/utils"
)

type mempool struct {
	Txs map[string]*Tx
}

var m *mempool
var memOnce sync.Once

func Mempool() *mempool {
	memOnce.Do(func() {
		if m == nil {
			m = &mempool{
				Txs: make(map[string]*Tx),
			}
		}
	})
	return m
}

func (m *mempool) AddTx(to string, amount int) {
	tx, err := makeTx(to, amount)
	utils.HandleErr(err)
	m.Txs[tx.Id] = tx
}