package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/erobx/trading-bot/pkg/app/model"
	_ "github.com/mattn/go-sqlite3"
)

const file string = "market.sqlite"
const testFile string = "test.sqlite"

const createSkinTable string = `
	CREATE TABLE IF NOT EXISTS skins (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	wear TEXT,
	price FLOAT
	);
`

const createStockTable string = `
	CREATE TABLE IF NOT EXISTS stocks (
	id TEXT NOT NULL PRIMARY KEY,
	price FLOAT,
	amount INTEGER
	);
`

// const createUserTable string = `
// 	CREATE TABLE IF NOT EXISTS users (
//  email TEXT,
// 	passwordHash TEXT,
// 	token TEXT,
// 	balance FLOAT
// 	);
// `

type Market struct {
	mu sync.RWMutex
	db *sql.DB
}

func NewMarket() (*Market, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	// db.Exec("DROP TABLE skins;")
	// db.Exec("DROP TABLE stocks;")
	if _, err = db.Exec(createSkinTable); err != nil {
		return nil, err
	}

	if _, err = db.Exec(createStockTable); err != nil {
		return nil, err
	}

	return &Market{
		db: db,
	}, nil
}

func NewTestMarket() (*Market, error) {
	db, err := sql.Open("sqlite3", testFile)
	if err != nil {
		return nil, err
	}
	db.Exec("DROP TABLE skins;")
	db.Exec("DROP TABLE stocks;")

	if _, err = db.Exec(createSkinTable); err != nil {
		return nil, err
	}

	if _, err = db.Exec(createStockTable); err != nil {
		return nil, err
	}

	return &Market{
		db: db,
	}, nil
}

func (m *Market) AddSkin(skin model.Skin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "INSERT INTO skins (id, name, wear, price) VALUES(NULL,?,?,?);"
	_, err := m.db.Exec(q, skin.Name, skin.Wear, skin.Price)
	if err != nil {
		return err
	}
	return nil
}

func (m *Market) GetSkin(name, wear string) (model.Skin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	skin := model.Skin{}
	q := "SELECT * FROM skins WHERE name=? AND wear=?"
	rows, err := m.db.Query(q, name, wear)
	if err != nil {
		return skin, false
	}
	defer rows.Close()

	return model.Skin{Name: name, Wear: wear}, true
}

func (m *Market) RemoveSkin(name, wear string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "DELETE FROM skins WHERE id IN (SELECT id FROM skins WHERE name=? AND wear=? LIMIT 1);"
	_, err := m.db.Exec(q, name, wear)

	return err == nil
}

func (m *Market) generateKey(name, wear string) string {
	return fmt.Sprintf("%s_%s", name, wear)
}
