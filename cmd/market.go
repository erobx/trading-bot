package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// DB
const file string = "market.sqlite"

const createSkinTable string = `
	CREATE TABLE IF NOT EXISTS skins (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	wear TEXT,
	price FLOAT
	);
`

const createUserTable string = `
	CREATE TABLE IF NOT EXISTS users (
    email TEXT,
	passwordHash TEXT
	);
`

// MARKET
type Market struct {
	mu sync.RWMutex
	db *sql.DB
}

func NewMarket() (*Market, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	db.Exec("DROP TABLE skins;")
	if _, err = db.Exec(createSkinTable); err != nil {
		return nil, err
	}

	return &Market{
		db: db,
	}, nil
}

func (m *Market) AddSkin(skin Skin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "INSERT INTO SKINS (id, name, wear, price) VALUES(NULL,?,?,?);"
	_, err := m.db.Exec(q, skin.Name, skin.Wear, skin.Price)
	if err != nil {
		return err
	}
	return nil
}

func (m *Market) GetSkin(name, wear string) (Skin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	skin := Skin{}
	q := "SELECT price FROM skins WHERE name=? AND wear=? ORDER BY price DESC"
	rows, err := m.db.Query(q, name, wear)
	if err != nil {
		return skin, false
	}
	defer rows.Close()

	var prices []float32
	for rows.Next() {
		var i float32
		err = rows.Scan(&i)
		if err != nil {
			return skin, false
		}
		prices = append(prices, i)
	}
	if len(prices) == 0 {
		return Skin{}, false
	}

	median:= getMedianPrice(prices)

	return Skin{Name: name, Wear: wear, Price: median}, true
}

func (m *Market) RemoveSkin(name, wear string, price float32) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return false
}

func (m *Market) generateKey(skin Skin) string {
	return fmt.Sprintf("%s_%s_%.2f", skin.Name, skin.Wear, skin.Price)
}

func getMedianPrice(prices []float32) float32  {
	return prices[len(prices)/2]
}

func randomPrices() []float32 {
	min := float32(23.10)
	max := float32(42.69)
	size := 10
	prices := make([]float32, size)
	for i := range prices {
		prices[i] = min + rand.Float32()*(max-min)
	}
	return prices
}
