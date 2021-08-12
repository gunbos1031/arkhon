package blockchain

import (
	"github.com/gunbos1031/arkhon/utils"
	"crypto/ecdsa"
	"os"
)

type wallet struct {
	privateKey 	*ecdsa.PrivateKey
	Address		string
	UtxOuts		map[string]*UtxOut
}

const (
	walletName = "cain.wallet"
)

var w *wallet

func (w *wallet) addUtxOut(tx *Tx) {
	for idx, txOut := range tx.TxOuts {
		if !isMine(w, txOut.Recipient) {
			// should treat another addr at future
			continue
		}
		uTxOut := &UtxOut{
			TxId: tx.Id, 
			Index: idx, 
			Amount: txOut.Amount,
		}
		uTxOut.getId()
		w.UtxOuts[uTxOut.Id] = uTxOut
	}	
}

func (w *wallet) setUtxOut(uTxos map[string]*UtxOut) {
	w.UtxOuts = uTxos
}

func (w *wallet) restore() {
	b := readFile()
	privKey := restorePrivKey(b)
	w.privateKey = privKey
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.restore()
		} else {
			key := generateKey()
			w.privateKey = key
			persist(w.privateKey)
		}
		w.Address = aFromKey(w.privateKey)
	}
	return w
}

func hasWalletFile() bool {
	_, err := os.Stat(walletName)
	return !os.IsNotExist(err)
}

func readFile() []byte {
	b, err := os.ReadFile(walletName)
	utils.HandleErr(err)
	return b
}

func persist(key *ecdsa.PrivateKey) {
	b := privToBytes(key)
	utils.HandleErr(os.WriteFile(walletName, b, 0644))
}

func getBalance() int {
	total := 0
	for _, uTxOut := range w.UtxOuts {
		total += uTxOut.Amount
	}
	return total
}

func isMine(w *wallet, addr string) bool {
	return w.Address == addr
}