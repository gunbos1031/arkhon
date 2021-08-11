package utils

import (
	"bytes"
	"log"
	"fmt"
	"crypto/sha256"
	"encoding/json"
	"strings"
	"math/big"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := json.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

func FromBytes(data []byte, i interface{}) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func Splitter(s, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r) - 1 < i {
		return ""
	}
	return r[i]
}

func EncodeBigInts(a, b *big.Int) string {
	z := append(a.Bytes(), b.Bytes()...)
	return fmt.Sprintf("%x", z)
}

func RestoreBigInts(payload string) (*big.Int, *big.Int, error){
	payloadAsBytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}
	firstHalves := payloadAsBytes[:len(payloadAsBytes)/2]
	secondHalves := payloadAsBytes[len(payloadAsBytes)/2:]
	a := &big.Int{}
	b := &big.Int{}
	a.SetBytes(firstHalves)
	b.SetBytes(secondHalves)
	return a, b, nil
}