package main

import (
	"github.com/gunbos1031/arkhon/db"
	"github.com/gunbos1031/arkhon/cli"
)

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}