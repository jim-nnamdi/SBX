package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func NewGenesis(p string) (Genesis, error) {
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return Genesis{}, err
	}
	var loadGenesis Genesis
	data := json.Unmarshal(content, &loadGenesis)
	if data != nil {
		log.Print("could not marshal data")
	}
	return loadGenesis, nil
}
