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

type walletStorage struct {
	PrivAsBytes 	[]byte
	UtxOutsAsBytes	[]byte
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
	persist(w)
}

func (w *wallet) setUtxOut(uTxos map[string]*UtxOut) {
	w.UtxOuts = uTxos
	persist(w)
}

func (w *wallet) restore() {
	storageAsBytes := readFile()
	var storage walletStorage
	var uTxos map[string]*UtxOut
	utils.FromBytes(storageAsBytes, &storage)
	privKey := restorePrivKey(storage.PrivAsBytes)
	utils.FromBytes(storage.UtxOutsAsBytes, &uTxos)
	w.privateKey = privKey
	w.UtxOuts = uTxos
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.restore()
		} else {
			key := generateKey()
			w.privateKey = key
			w.UtxOuts = make(map[string]*UtxOut)
			persist(w)
		}
		w.Address = aFromKey(w.privateKey)
	}
	return w
}

func persist(wallet *wallet) {
	privAsBytes := privToBytes(wallet.privateKey)
	uTxoAsbytes := utils.ToBytes(wallet.UtxOuts)
	storage := walletStorage{privAsBytes, uTxoAsbytes}
	storageAsBytes := utils.ToBytes(storage)
	utils.HandleErr(os.WriteFile(walletName, storageAsBytes, 0644))
}

func hasWalletFile() bool {
	_, err := os.Stat(walletName)
	return !os.IsNotExist(err)
}

func readFile() []byte {
	storageAsBytes, err := os.ReadFile(walletName)
	utils.HandleErr(err)
	return storageAsBytes
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