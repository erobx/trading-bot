package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// DB
const file string = "skins.sqlite"

const create string = `
	CREATE TABLE IF NOT EXISTS skins (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	wear TEXT,
	price FLOAT
	);
`

// MARKET
type Market struct {
	mu    sync.RWMutex
	Skins map[string]map[Skin]int
	db    *sql.DB
}

func NewMarket() (*Market, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(create); err != nil {
		return nil, err
	}

	return &Market{
		Skins: make(map[string]map[Skin]int),
		db:    db,
	}, nil
}

func (m *Market) AddSkin(skin Skin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Map
	key := m.generateKey(skin)
	if m.Skins[key] == nil {
		m.Skins[key] = make(map[Skin]int)
	}
	m.Skins[key][skin]++

	// DB
	q := "INSERT INTO SKINS (id, name, wear, price) VALUES(NULL,?,?,?);"
	_, err := m.db.Exec(q, skin.Name, skin.Wear, skin.Price)
	if err != nil {
		return err
	}
	return nil
}

func (m *Market) GetSkin(name, wear string, price float32) (Skin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := m.generateKey(NewSkin(name, wear, price))
	if skinsMap, ok := m.Skins[key]; ok {
		for skin := range skinsMap {
			return skin, true
		}
	}
	return Skin{}, false
}

func (m *Market) RemoveSkin(name, wear string, price float32) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := m.generateKey(NewSkin(name, wear, price))
	if skinsMap, ok := m.Skins[key]; ok {
		for skin := range skinsMap {
			m.Skins[key][skin]--
			if m.Skins[key][skin] == 0 {
				delete(m.Skins[key], skin)
			}
			return true
		}
	}
	return false
}

func (m *Market) generateKey(skin Skin) string {
	return fmt.Sprintf("%s_%s_%.2f", skin.Name, skin.Wear, skin.Price)
}
