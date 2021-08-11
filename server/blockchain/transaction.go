package blockchain

import (

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

func addTx(to string, amount int) {
	from := Wallet().Address
	balance := getBalance()
	// make TxIns
}