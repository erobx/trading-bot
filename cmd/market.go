package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
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
	passwordHash TEXT,
	token TEXT,
	balance FLOAT
	);
`

// Price BS
type dbDecimal decimal.Decimal

func (v *dbDecimal) Scan(value interface{}) error {
	if value == nil {
		*v = dbDecimal(decimal.Zero)
		return nil
	}
	if sv, err := driver.String.ConvertValue(value); err == nil {
		if vv, ok := sv.(string); ok {
			if vvv, err := decimal.NewFromString(vv); err == nil {
				*v = dbDecimal(vvv)
				return nil
			}
		}
	}
	return errors.New("cannot convert to decimal")
}

func (v dbDecimal) Value() (driver.Value, error) {
	dec := decimal.Decimal(v)
	price, _ := dec.Float64()
	return price, nil
}

func (v dbDecimal) String() string {
	return decimal.Decimal(v).String()
}

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
	// db.Exec("DROP TABLE skins;")
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

	var prices []dbDecimal
	for rows.Next() {
		var i dbDecimal
		err = rows.Scan(&i)
		if err != nil {
			return skin, false
		}
		prices = append(prices, i)
	}
	if len(prices) == 0 {
		return Skin{}, false
	}

	median := getMedianPrice(prices)

	return Skin{Name: name, Wear: wear, Price: median}, true
}

func (m *Market) RemoveSkin(name, wear string, price dbDecimal) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return false
}

func (m *Market) generateKey(skin Skin) string {
	return fmt.Sprintf("%s_%s_%.2f", skin.Name, skin.Wear, skin.Price)
}

func getMedianPrice(prices []dbDecimal) dbDecimal {
	return prices[len(prices)/2]
}

func randomPrices() []dbDecimal {
	min_d := float64(23.12)
	max_d := float64(42.99)
	size := 10
	prices := make([]dbDecimal, size)

	for i := range prices {
		d := min_d + rand.Float64() * max_d
		prices[i] = dbDecimal(decimal.NewFromFloat(d))
	}

	return prices
}
