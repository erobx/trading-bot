package db

import (
	"database/sql"
	"log"
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
	price FLOAT,
	gun TEXT,
	min FLOAT,
	max FLOAT
	);
`

const createGroupTable string = `
	CREATE TABLE IF NOT EXISTS groups (
	id INTEGER NOT NULL PRIMARY KEY,
	tier INTEGER,
	active INTEGER
	);
`

const createGroupSkinTable string = `
	CREATE TABLE IF NOT EXISTS group_skins (
	id INTEGER NOT NULL PRIMARY KEY,
	group_id INTEGER NOT NULL,
	skin_id INTEGER NOT NULL,
	FOREIGN KEY(group_id) REFERENCES groups(id),
	FOREIGN KEY(skin_id) REFERENCES skins(id)
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

/*
M-to-M joining
SELECT g.id AS group_id, g.tier, gs.skin_id, s.name
  FROM groups g
JOIN group_skins gs ON (g.id = gs.group_id)
JOIN skins s ON (gs.skin_id = s.id);
*/

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
	//Db.Exec("DROP TABLE skins;")
	//Db.Exec("DROP TABLE groups;")
	//Db.Exec("DROP TABLE group_skins;")
	if _, err = Db.Exec(createSkinTable); err != nil {
		return nil, err
	}

	if _, err = Db.Exec(createGroupTable); err != nil {
		return nil, err
	}

	if _, err = Db.Exec(createGroupSkinTable); err != nil {
		return nil, err
	}

	return &Market{
		Db: Db,
	}, nil
}

func (m *Market) AddSkin(skin model.Skin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "INSERT INTO SKINS (id, name, wear, price, gun, min, max) VALUES(NULL,?,?,?,?,?,?);"
	_, err := m.Db.Exec(q, skin.Name, skin.Wear, skin.Price, skin.Gun, skin.Min, skin.Max)
	if err != nil {
		return err
	}
	return nil
}

func (m *Market) AddSkinToGroup(gid string, sid string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "INSERT INTO GROUP_SKINS (id, group_id, skin_id) VALUES(NULL,?,?);"
	_, err := m.Db.Exec(q, gid, sid)
	if err != nil {
		return err
	}
	return nil
}

func (m *Market) GetSkin(sid string) (model.Skin, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	skin := model.Skin{}
	q := "SELECT name,wear,price,gun,min,max FROM skins WHERE id=?"
	rows, err := m.Db.Query(q, sid)
	if err != nil {
		return skin, err
	}
	defer rows.Close()

	err = rows.Scan(&skin.Name, &skin.Wear, &skin.Price, &skin.Gun, &skin.Min, &skin.Max)

	return skin, err
}

func (m *Market) AddGroup(group model.Group) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := "INSERT INTO GROUPS (id, tier, active) VALUES(NULL,?,?);"
	_, err := m.Db.Exec(q, group.Tier, group.Active)
	if err != nil {
		return err
	}
	return nil
}

func (m *Market) GetActiveGroups() ([]model.DisplayGroup, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	q := `
	SELECT g.id AS group_id, g.tier, s.name, s.wear, s.price, s.gun,
		s.min, s.max
		FROM groups g
	  JOIN group_skins gs ON (g.id=gs.group_id)
	  JOIN skins s ON (gs.skin_id=s.id)
	  ORDER BY g.id, s.id;
	`
	rows, err := m.Db.Query(q)
	if err != nil {
		return []model.DisplayGroup{}, err
	}
	defer rows.Close()

	var groups []model.DisplayGroup
	var currentGroup *model.DisplayGroup

	for rows.Next() {
		var gID int
		var gTier string
		var sName string
		var sWear string
		var sPrice types.DbDecimal
		var sGun string
		var sMin types.DbDecimal
		var sMax types.DbDecimal

		err := rows.Scan(&gID, &gTier, &sName, &sWear, &sPrice, &sGun, &sMin, &sMax)
		if err != nil {
			return groups, err
		}

		if currentGroup == nil || currentGroup.GroupId != gID {
			if currentGroup != nil {
				groups = append(groups, *currentGroup)
			}
			currentGroup = &model.DisplayGroup{
				GroupId: gID,
				Tier:    gTier,
				Skins:   []model.Skin{},
			}
		}

		sTemp := model.Skin{
			Name:  sName,
			Wear:  sWear,
			Price: sPrice,
			Gun:   sGun,
			Min:   sMin,
			Max:   sMax,
		}
		currentGroup.Skins = append(currentGroup.Skins, sTemp)
	}
	rows.Close()

	if currentGroup != nil {
		groups = append(groups, *currentGroup)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return groups, nil
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
