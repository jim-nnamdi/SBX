package database

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
)

type Hash [32]byte

func (h *Hash) marshal() ([]byte, error) {
	return []byte(hex.EncodeToString(h[:])), nil
}

func (h *Hash) Unmarshal(data []byte) error {
	_, err := hex.Decode(h[:], data)
	return err
}

type Block struct {
	Header BlockHeader `json:"header"`
	Tx     []Tx        `json:"transactions"`
}

type BlockHeader struct {
	Parent Hash `json:"parent"`
	Time   uint `json:"time"`
}

type BlockFS struct {
	Key   Hash  `json:"key"`
	Value Block `json:"value"`
}

func (b *Block) Hash() (Hash, error) {
	blockjson, err := json.Marshal(b)
	if err != nil {
		log.Print("Could not marshal json properly")
	}
	return sha256.Sum256(blockjson), nil
}
