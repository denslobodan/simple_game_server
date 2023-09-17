package db

import (
	"app/server"
	"sync"
)

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	store map[string]int
	//  A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.store[name]
}

// RecordWin will record a player's win.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

// GetPlayersScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetLeague() []server.Player {
	var league []server.Player
	for name, wins := range i.store {
		league = append(league, server.Player{Name: name, Wins: wins})
	}
	return league
}
