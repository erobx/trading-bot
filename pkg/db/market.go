package db

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/types"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
)

// Db
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
	passwordHash TEXT,
	token TEXT,
	balance FLOAT
	);
`

// MARKET
type Market struct {
	mu sync.RWMutex
	Db *sql.DB
}

func NewMarket() (*Market, error) {
	Db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	// Db.Exec("DROP TABLE skins;")
	if _, err = Db.Exec(createSkinTable); err != nil {
		return nil, err
	}

	return &Market{
		Db: Db,
	}, nil
}

func (m *Market) AddSkin(skin model.Skin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "INSERT INTO SKINS (id, name, wear, price) VALUES(NULL,?,?,?);"
	_, err := m.Db.Exec(q, skin.Name, skin.Wear, skin.Price)
	if err != nil {
		return err
	}
	return nil
}

func (m *Market) GetSkin(name, wear string) (model.Skin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	skin := model.Skin{}
	q := "SELECT price FROM skins WHERE name=? AND wear=? ORDER BY price DESC"
	rows, err := m.Db.Query(q, name, wear)
	if err != nil {
		return skin, false
	}
	defer rows.Close()

	var prices []types.DbDecimal
	for rows.Next() {
		var i types.DbDecimal
		err = rows.Scan(&i)
		if err != nil {
			return skin, false
		}
		prices = append(prices, i)
	}
	if len(prices) == 0 {
		return model.Skin{}, false
	}

	median := getMedianPrice(prices)

	return model.Skin{Name: name, Wear: wear, Price: median}, true
}

func (m *Market) RemoveSkin(name, wear string, price types.DbDecimal) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return false
}

func (m *Market) generateKey(skin model.Skin) string {
	return fmt.Sprintf("%s_%s_%.2f", skin.Name, skin.Wear, skin.Price)
}

func getMedianPrice(prices []types.DbDecimal) types.DbDecimal {
	return prices[len(prices)/2]
}

func RandomPrices() []types.DbDecimal {
	min_d := float64(23.12)
	max_d := float64(42.99)
	size := 10
	prices := make([]types.DbDecimal, size)

	for i := range prices {
		d := min_d + rand.Float64()*max_d
		prices[i] = types.DbDecimal(decimal.NewFromFloat(d))
	}

	return prices
}
