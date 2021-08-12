package blockchain

import (
	"github.com/gunbos1031/arkhon/utils"
	"encoding/hex"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"time"
)

type Tx struct {
	Id			string		`json:"id"`
	Timestamp	int			`json:"timestamp"`
	Sender		string		`json:"sender"`
	Recipient	string		`json:"recipient"`
	TxIns		[]*TxIn		`json:"txIns"`
	TxOuts		[]*TxOut	`json:"txOuts"`
	Signature	string		`json:"signature"`
}

type TxIn struct {
	TxId		string		`json:"txId"`
	Index		int			`json:"index"`
	Signature	string		`json:"signature"`
}

type TxOut struct {
	Recipient	string		`json:"recipient"`
	Amount		int			`json:"amount"`
}

type UtxOut struct {
	Id			string		`json:"id"`
	TxId		string		`json:"txId"`
	Index		int			`json:"index"`
	Amount		int			`json:"amount"`
}

var (
	ErrNoMoney = errors.New("Not enough money")
	ErrInvalid = errors.New("Invalid transaction")
)

const mineReward int = 50

func (t *Tx) getId() {
	t.Timestamp = int(time.Now().Unix())
	hash := utils.Hash(t)
	t.Id = hash
}

func (t *Tx) sign(wallet *wallet) {
	txIdAsBytes, err := hex.DecodeString(t.Id)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, wallet.privateKey, txIdAsBytes)
	utils.HandleErr(err)
	signature := utils.EncodeBigInts(r, s)
	t.Signature = signature
	for _, txIn := range t.TxIns {
		txIn.Signature = signature
	}
}

func (u *UtxOut) getId() {
	hash := utils.Hash(u)
	u.Id = hash
}

func makeTx(to string, amount int) (*Tx, error) {
	balance := getBalance()
	if balance < amount {
		return nil, ErrNoMoney
	}
	
	var txIns []*TxIn
	var txOuts []*TxOut
	wallet := Wallet()
	total := 0
	from := wallet.Address
	uTxOuts := wallet.UtxOuts
	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{
			TxId: uTxOut.TxId,
			Index: uTxOut.Index,
			Signature: "",
		}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
		delete(uTxOuts, uTxOut.Id)
	}
	wallet.setUtxOut(uTxOuts)
	if change := total - amount; change != 0 {
		txOuts = append(txOuts, &TxOut{from, change})
	}
	txOuts = append(txOuts, &TxOut{to, amount})
	tx := &Tx{
		Sender: from,
		Recipient: to,
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()
	tx.sign(wallet)
	if !verify(tx) {
		return nil, ErrInvalid
	}
	wallet.addUtxOut(tx)
	return tx, nil
}

func verify(t *Tx) bool {
	valid := true
	for _, txIn := range t.TxIns {
		prevTx := findTx(Blockchain(), txIn.TxId)
		if prevTx == nil {
			valid = false
			break
		}
		addr := prevTx.TxOuts[txIn.Index].Recipient
		valid = doVerify(txIn.Signature, t.Id, addr)
		if !valid {
			break
		}
	}
	return valid
}

func makeCoinbaseTx() *Tx {
	txIn := &TxIn{"", -1, ""}
	txOut := &TxOut{Wallet().Address, mineReward}
	tx := &Tx{
		Sender: "COINBASE",
		Recipient: Wallet().Address,
		TxIns: []*TxIn{txIn},
		TxOuts: []*TxOut{txOut},
	}
	tx.getId()
	return tx
}