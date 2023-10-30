package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func main() {

}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	//[][]byte{} expression is used to create a slice of byte slices
	//[]byte{} as the separator, which means there will be no separator between the concatenated slices

	hash := sha256.Sum256(headers)

	b.Hash = hash[:]

}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block(time.Now().Unix(), []byte(data), prevBlockHash, []byte{})
	block.SetHash()
	return block
}
