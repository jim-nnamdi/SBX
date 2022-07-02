package database

import (
	"log"
	"os"
)

type State struct {
	Balances        map[Account]uint `json:"balances"`
	TxMempool       []Tx             `json:"txMempool"`
	DbFile          *os.File         `json:"dbfile"`
	LatestBlockHash Hash             `json:"latestblockhash"`
}

func (s *State) ApplyBlock(b Block) error {
	for _, tx := range b.Tx {
		if err := s.Apply(tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *State) Apply(tx Tx) error {
	if tx.IsReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}
	if s.Balances[tx.From] < tx.Value {
		log.Print("insufficient balances")
	}
	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value
	return nil
}
