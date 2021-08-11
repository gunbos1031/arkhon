package blockchain

import (
	"github.com/gunbos1031/arkhon/utils"
	"crypto/ecdsa"
	"os"
	"fmt"
)

type wallet struct {
	privateKey 	*ecdsa.PrivateKey
	Address		string
}

const (
	walletName = "cain.wallet"
)

var w *wallet

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
	fmt.Println(w.Address)
	return w
}

func (w *wallet) restore() {
	b := readFile()
	privKey := restorePrivKey(b)
	w.privateKey = privKey
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