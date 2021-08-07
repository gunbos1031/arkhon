package db

import (
	bolt "go.etcd.io/bbolt"
	"github.com/gunbos1031/arkhon/utils"
)

const (
	// dbName = "blockchain"
	dataBucket = "data"
	blocksBucket = "blocks"
	checkpoint = "checkpoint"
)

var db *bolt.DB

func InitDB() {
	if db == nil {
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil)
		utils.HandleErr(err)
		db = dbPointer
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
}

func Close() {
	db.Close()
}

func SaveBlockchain(data []byte) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlock(key, value []byte) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put(key, value)
		return err
	})
	utils.HandleErr(err)
}

func LoadBlockchain() []byte {
	var data []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	utils.HandleErr(err)
	return data
}

func FindBlock(hash string) []byte {
	var data []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))	
		return nil
	})
	utils.HandleErr(err)
	return data
}