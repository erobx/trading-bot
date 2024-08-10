package db

import (
	"database/sql"
	"log"
	"sync"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/types"
	_ "github.com/mattn/go-sqlite3"
)

const file string = "market.sqlite"

const createSkinsTable string = `
	CREATE TABLE IF NOT EXISTS skins (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	weapon TEXT,
	wear TEXT,
	color TEXT,
	collection TEXT,
	float_min FLOAT,
	float_max FLOAT
	);
`

const createTradeupsTable string = `
	CREATE TABLE IF NOT EXISTS tradeups (
	id INTEGER NOT NULL PRIMARY KEY,
	tier INTEGER,
	active INTEGER
	);
`

const createUsersTable string = `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY,
	username TEXT,
	email TEXT,
	hash TEXT
	);
`

const createTradeupSkinsTable string = `
	CREATE TABLE IF NOT EXISTS tradeup_skins (
	id INTEGER NOT NULL PRIMARY KEY,
	status TEXT,
	tradeup_id INTEGER NOT NULL,
	user_inv_id INTEGER NOT NULL,

	FOREIGN KEY(tradeup_id) REFERENCES tradeups(id),
	FOREIGN KEY(user_inv_id) REFERENCES user_inventory(id),
	UNIQUE(tradeup_id, user_inv_id)
	);
`

const createUserInventoryTable string = `
	CREATE TABLE IF NOT EXISTS user_inventory (
	id INTEGER NOT NULL PRIMARY KEY,
	fl FLOAT,
	price FLOAT,
	user_id INTEGER NOT NULL,
	skin_id INTEGER NOT NULL,

	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(skin_id) REFERENCES skins(id)
	);
`

var tables = []string{createSkinsTable, createTradeupsTable, createUsersTable, createTradeupSkinsTable, createUserInventoryTable}

type Db struct {
	mu    sync.RWMutex
	sqlDb *sql.DB
}

func createTables(db *sql.DB) {
	for _, t := range tables {
		if _, err := db.Exec(t); err != nil {
			panic(err)
		}
	}
}

func NewDbConn() (*Db, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	createTables(db)

	return &Db{
		sqlDb: db,
	}, nil
}

func (d *Db) AddUser(username, email, hash string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	q := "INSERT INTO users (id,username,email,hash) VALUES(NULL,?,?,?)"
	_, err := d.sqlDb.Exec(q, username, email, hash)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) AddSkin(skin model.Skin) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	q := "INSERT INTO skins (id, name, weapon, wear, color, collection, float_min, float_max) VALUES(NULL,?,?,?,?,?,?,?);"
	_, err := d.sqlDb.Exec(q, skin.Name, skin.Weapon, skin.Wear, skin.Color, skin.Collection, skin.FloatMin, skin.FloatMax)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) AddSkinToTradeup(tid string, sid string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	q := "INSERT INTO tradeup_skins (id, tradeup_id, user_skin_id) VALUES(NULL,?,?);"
	_, err := d.sqlDb.Exec(q, tid, sid)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) GetInventory(uid string) ([]model.Skin, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	q := `
		SELECT u.id AS user_id, ui.fl, s.name, s.wear
			FROM users u
		JOIN user_inventory ui ON (u.id=ui.user_id)
		JOIN skins s ON (ui.skin_id=s.id)
			WHERE u.id=?;
	`

	var skins []model.Skin

	rows, err := d.sqlDb.Query(q, uid)
	if err != nil {
		return skins, err
	}

	for rows.Next() {
		var uid string
		sTemp := model.Skin{}

		err := rows.Scan(&uid, &sTemp.Fl, &sTemp.Name, &sTemp.Wear)
		if err != nil {
			return skins, err
		}

		skins = append(skins, sTemp)
	}
	rows.Close()

	return skins, nil
}

func (d *Db) AddTradeup(group model.Tradeup) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	q := "INSERT INTO tradeups (id, tier, active) VALUES(NULL,?,?);"
	_, err := d.sqlDb.Exec(q, group.Tier, group.Active)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) GetActiveTradeups() ([]model.DisplayTrade, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	q := `
	SELECT t.id AS tradeup_id, t.tier, s.name, s.weapon, s.wear, s.color, s.collection, s.float_min, s.float_max
		FROM tradeups t
	JOIN tradeup_skins ts ON (t.id=ts.tradeup_id)
	JOIN skins s ON (ts.user_inv_id=s.id)
		ORDER BY t.id, s.id;
	`
	rows, err := d.sqlDb.Query(q)
	if err != nil {
		return []model.DisplayTrade{}, err
	}
	defer rows.Close()

	var tradeups []model.DisplayTrade
	var currentTradeup *model.DisplayTrade

	for rows.Next() {
		var tid int
		var tier string
		var name string
		var weapon string
		var wear string
		var color string
		var collection string
		var floatMin types.DbDecimal
		var floatMax types.DbDecimal

		err := rows.Scan(&tid, &tier, &name, &weapon, &wear, &color, &collection, &floatMin, &floatMax)
		if err != nil {
			return tradeups, err
		}

		if currentTradeup == nil || currentTradeup.TradeId != tid {
			if currentTradeup != nil {
				tradeups = append(tradeups, *currentTradeup)
			}
			currentTradeup = &model.DisplayTrade{
				TradeId: tid,
				Tier:    tier,
				Skins:   []model.Skin{},
			}
		}

		sTemp := model.Skin{
			Name:     name,
			Wear:     wear,
			FloatMin: floatMin,
			FloatMax: floatMax,
		}
		currentTradeup.Skins = append(currentTradeup.Skins, sTemp)
	}
	rows.Close()

	if currentTradeup != nil {
		tradeups = append(tradeups, *currentTradeup)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return tradeups, nil
}

func (d *Db) GetChangedTradeup(tid string) (model.DisplayTrade, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	q := `
		SELECT t.id AS tradeup_id, t.tier, s.name, s.weapon, s.wear, s.color, s.collection, s.float_min, s.float_max
			FROM tradeups t
		JOIN tradeup_skins ts ON (t.id=ts.tradeup_id)
		JOIN skins s ON (ts.user_inv_id=s.id)
			WHERE t.id=?;
	`
	rows, err := d.sqlDb.Query(q, tid)
	if err != nil {
		return model.DisplayTrade{}, err
	}
	defer rows.Close()

	var tradeup model.DisplayTrade
	var id int
	var tier string

	tradeup = model.DisplayTrade{
		Skins: []model.Skin{},
	}

	for rows.Next() {
		var sName string
		var sWeapon string
		var sWear string
		var sMin types.DbDecimal
		var sMax types.DbDecimal

		err := rows.Scan(&id, &tier, &sName, &sWeapon, &sWear, &sMin, &sMax)
		if err != nil {
			return tradeup, err
		}

		sTemp := model.Skin{
			Name:     sName,
			Wear:     sWear,
			FloatMin: sMin,
			FloatMax: sMax,
		}
		tradeup.Skins = append(tradeup.Skins, sTemp)
	}
	rows.Close()

	tradeup.TradeId = id
	tradeup.Tier = tier

	return tradeup, nil
}
