package blockchain

import (
	"github.com/gunbos1031/arkhon/utils"
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
	TxId		string		`json:"txId"`
	Index		int			`json:"index"`
	Amount		int			`json:"amount"`
}

var (
	ErrNoMoney = errors.New("Not enough money")
	ErrInvalid = errors.New("Invalid transaction")
)

func (t *Tx) getId() {
	t.Timestamp = int(time.Now().Unix())
	hash := utils.Hash(t)
	t.Id = hash
}

func AddTx(to string, amount int) {
	tx, err := makeTx(Wallet(), to, amount)
}

func makeTx(wallet *wallet, to string, amount int) (*Tx, error) {
	balance := getBalance()
	if balance < amount {
		return nil, ErrNoMoney
	}
	
	var txIns []*TxIn
	var txOuts []*TxOut
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
	}
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
		addr := prevTx.TxOuts[txIn.Index].Address
		valid = doVerify(txIn.Signature, tx.Id, addr)
		if !valid {
			break
		}
	}
	return valid
}