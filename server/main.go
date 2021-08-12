package main

import (
	"github.com/gunbos1031/arkhon/db"
	"github.com/gunbos1031/arkhon/cli"
	"github.com/gunbos1031/arkhon/blockchain"
)

func main() {
	defer blockchain.Persist(blockchain.Wallet())
	defer db.Close()
	db.InitDB()
	cli.Start()
}