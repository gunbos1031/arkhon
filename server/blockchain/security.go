package blockchain

import (
	"github.com/gunbos1031/arkhon/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
)

func doVerify(signature, id, addr string) bool {
	r, s, err := utils.RestoreBigInts(signature)
	utils.HandleErr(err)
	x, y, err := utils.RestoreBigInts(addr)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X: x,
		Y: y,
	}
	idAsBytes, err := hex.DecodeString(id)
	utils.HandleErr(err)
	ok := ecdsa.Verify(&publicKey, idAsBytes, r, s)
	return ok
}

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