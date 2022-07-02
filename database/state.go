package database

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type State struct {
	Balances        map[Account]uint `json:"balances"`
	TxMempool       []Tx             `json:"txMempool"`
	DbFile          *os.File         `json:"dbfile"`
	LatestBlockHash Hash             `json:"latestblockhash"`
}

func (s *State) NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	gen, err := NewGenesis(filepath.Join(cwd, "database", "genesis.json"))
	if err != nil {
		log.Print(err)
	}
	balances := make(map[Account]uint, 0)
	for acc, bal := range gen.Balances {
		balances[acc] = bal
	}
	f, err := os.OpenFile(filepath.Join(cwd, "database", "block.db"), os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		log.Print(err)
	}
	scanner := bufio.NewScanner(f)
	state := &State{Balances: balances, TxMempool: make([]Tx, 0), DbFile: f, LatestBlockHash: Hash{}}
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Print(err)
		}
		blockFsJson := scanner.Bytes()
		var blockfs BlockFS
		err := json.Unmarshal(blockFsJson, &blockfs)
		if err != nil {
			log.Print(err)
		}
		err = state.ApplyBlock(blockfs.Value)
		if err != nil {
			log.Print(err)
		}
		state.LatestBlockHash = blockfs.Key
	}
	return state, nil
}

func (s *State) Persist() {

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
