package main

import (
	"github.com/gunbos1031/arkhon/rest"
	"github.com/gunbos1031/arkhon/db"
)

func main() {
	defer db.Close()
	db.InitDB()
	rest.Start()
}