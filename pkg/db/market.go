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
//     email TEXT,
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
	q := "SELECT price FROM skins WHERE name=? AND wear=? ORDER BY price DESC"
	rows, err := m.db.Query(q, name, wear)
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

func (m *Market) RemoveSkin(name, wear string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "DELETE FROM skins WHERE id IN (SELECT id FROM skins WHERE name=? AND wear=? LIMIT 1);"
	_, err := m.db.Exec(q, name, wear)

	return err == nil
}

func (m *Market) GetStock(name, wear string) (model.Stock, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := m.generateKey(name, wear)
	stock, ok := m.stockExists(key)
	if !ok {
		return stock, false
	}

	return stock, true
}

func (m *Market) UpdateStock(skin model.Skin, add bool) bool {
	key := m.generateKey(skin.Name, skin.Wear)
	stock, ok := m.stockExists(key)

	if !ok && !add {
		return false
	}

	if add {
		if !ok {
			m.mu.Lock()
			defer m.mu.Unlock()
			q := "INSERT INTO stocks (id, price, amount) VALUES(?,?,?);"
			_, err := m.db.Exec(q, key, skin.Price, 1)
			return err == nil
		}

		newPrice := m.getNewStockPrice(skin.Name, skin.Wear)
		m.mu.Lock()
		defer m.mu.Unlock()

		q := "UPDATE stocks SET price=?, amount=? WHERE id=?;"
		_, err := m.db.Exec(q, newPrice, stock.Amount+1, key)

		return err == nil
	}

	newPrice := m.getNewStockPrice(skin.Name, skin.Wear)
	m.mu.Lock()
	defer m.mu.Unlock()

	if stock.Amount == 1 {
		q := "DELETE FROM stocks WHERE id=?;"
		_, err := m.db.Exec(q, key)
		return err == nil
	}

	q := "UPDATE stocks SET price=?, amount=? WHERE id=?;"
	_, err := m.db.Exec(q, newPrice, stock.Amount-1, key)

	return err == nil
}

func (m *Market) stockExists(value string) (model.Stock, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	stock := model.Stock{}
	q := "SELECT * FROM stocks WHERE id='" + value + "';"
	row := m.db.QueryRow(q)

	err := row.Scan(&stock.Id, &stock.Price, &stock.Amount)
	if err != nil {
		return stock, false
	}

	return stock, true
}

func (m *Market) getNewStockPrice(name, wear string) types.DbDecimal {
	m.mu.RLock()
	defer m.mu.RUnlock()

	q := "SELECT price FROM skins WHERE name=? AND wear=? ORDER BY price DESC"
	rows, err := m.db.Query(q, name, wear)
	if err != nil {
		return types.DbDecimal{}
	}
	defer rows.Close()

	var prices []types.DbDecimal
	for rows.Next() {
		var i types.DbDecimal
		err = rows.Scan(&i)
		if err != nil {
			return types.DbDecimal{}
		}
		prices = append(prices, i)
	}
	if len(prices) == 0 {
		return types.DbDecimal{}
	}

	return getMedianPrice(prices)
}

func getMedianPrice(prices []types.DbDecimal) types.DbDecimal {
	return prices[len(prices)/2]
}

func (m *Market) generateKey(name, wear string) string {
	return fmt.Sprintf("%s_%s", name, wear)
}

func RandomPrices() []types.DbDecimal {
	min_d := float64(23.12)
	max_d := float64(42.99)
	size := 10
	prices := make([]types.DbDecimal, size)

	for i := range prices {
		d := min_d + rand.Float64()*max_d
		s := fmt.Sprintf("%.2f", d)
		temp, _ := decimal.NewFromString(s)
		prices[i] = types.DbDecimal(temp)
	}

	return prices
}
