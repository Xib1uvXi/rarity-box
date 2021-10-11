package storage

import (
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/btcsuite/goleveldb/leveldb/errors"
)

var MemDaoNotFound = errors.New("Not Found")

type memDao struct {
	db map[string][]*types.Summoner
}

func (m *memDao) Save(address string, summoners []*types.Summoner) error {
	m.db[address] = summoners
	return nil
}

func (m *memDao) SummonersByAddress(address string) ([]*types.Summoner, error) {
	summoners, ok := m.db[address]

	if !ok {
		return nil, MemDaoNotFound
	}

	return summoners, nil
}
