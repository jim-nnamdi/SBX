package database

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

type Genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func NewGenesis(p io.Reader) Genesis {
	content, err := ioutil.ReadAll(p)
	if err != nil {
		return Genesis{}
	}
	var loadGenesis Genesis
	data := json.Unmarshal(content, &loadGenesis)
	if data != nil {
		log.Print("could not marshal data")
	}
	return loadGenesis
}
