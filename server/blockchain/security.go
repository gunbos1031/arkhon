package blockchain

import (
	"github.com/gunbos1031/arkhon/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
)

func generateKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func privToBytes(privKey *ecdsa.PrivateKey) []byte {
	b, err := x509.MarshalECPrivateKey(privKey)
	utils.HandleErr(err)
	return b
}

func restorePrivKey(privAsBytes []byte) *ecdsa.PrivateKey {
	privKey, err := x509.ParseECPrivateKey(privAsBytes)
	utils.HandleErr(err)
	return privKey
}

func aFromKey(key *ecdsa.PrivateKey) string {
	return utils.EncodeBigInts(key.X, key.Y)
}