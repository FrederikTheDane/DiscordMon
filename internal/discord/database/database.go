package database

import (
	"database/sql"
	"fmt"
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
	_ "github.com/lib/pq"
	"os"
)

var database *sql.DB

//This must be called before using the database
func Connect() error {
	var err error

	dbHost := os.Getenv("DISCORDMON_DB_HOST")
	dbPort := os.Getenv("DISCORDMON_DB_PORT")
	dbUser := os.Getenv("DISCORDMON_DB_USER")
	dbPwd := os.Getenv("DISCORDMON_DB_PWD")
	dbName := os.Getenv("DISCORDMON_DB_NAME")
	conStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPwd, dbHost, dbPort, dbName)
	database, err = sql.Open("postgres", conStr)

	return err
}

func AddNewPokemon(id int, name string, gendered bool, types, hp, attack, defense, sattack, sdefense, speed int) (*mechanics.PokeMonBase, error) {
	_, err := database.Exec(`
INSERT INTO base_monsters (monster_id, gendered, types, base_hp, base_attack, base_defense, base_sattack, base_sdefense, base_speed, name) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
`, id, gendered, types, hp, attack, defense, sattack, sdefense, speed, name)

	if err != nil {
		return nil, err
	}

	return GetMonBaseByID(int64(id))
}

func GetMonBaseByID(id int64) (*mechanics.PokeMonBase, error) {
	rows, err := database.Query(`
SELECT monster_id, gendered, types, base_hp, base_attack, base_defense, base_sattack, base_sdefense, base_speed, name FROM base_monsters WHERE monster_id = $1;
`, id)
	if err != nil {
		return nil, err
	}

	monid, types, hp, attack, defense, sattack, sdefense, speed := new(int), new(int), new(int), new(int), new(int), new(int), new(int), new(int)
	gendered := new(bool)
	name := new(string)

	rows.Next()
	err = rows.Scan(monid, gendered, types, hp, attack, defense, sattack, sdefense, speed, name)
	if err != nil {
		return nil, err
	}

	mon := mechanics.PokeMonBase{
		Gendered:  *gendered,
		ID:        *monid,
		Types:     constants.GetTypesFromID(*monid),
		BaseStats: [6]int{*hp, *attack, *defense, *sattack, *sdefense, *speed},
		Name:      *name,
	}

	return &mon, err
}

func SetupTables() error {
	var err error

	_, err = database.Exec(`
CREATE TABLE IF NOT EXISTS public.base_monsters (
    id				BIGSERIAL PRIMARY KEY,
    monster_id		BIGINT NOT NULL UNIQUE,
    gendered		BOOLEAN NOT NULL,
    types			INTEGER NOT NULL,
    
    base_hp			SMALLINT NOT NULL,
    base_attack		SMALLINT NOT NULL,
    base_defense	SMALLINT NOT NULL,
    base_sattack	SMALLINT NOT NULL,
    base_sdefense	SMALLINT NOT NULL,
    base_speed		SMALLINT NOT NULL,
    
    yield_hp		SMALLINT NOT NULL DEFAULT 0,
    yield_attack	SMALLINT NOT NULL DEFAULT 0,
    yield_defense	SMALLINT NOT NULL DEFAULT 0,
    yield_sattack	SMALLINT NOT NULL DEFAULT 0,
    yield_sdefense	SMALLINT NOT NULL DEFAULT 0,
    yield_speed		SMALLINT NOT NULL DEFAULT 0,
    
    name			TEXT NOT NULL UNIQUE
);
	`)

	if err != nil {
		return err
	}

	_, err = database.Exec(`
CREATE TABLE IF NOT EXISTS public.monster_types (
    id			SERIAL PRIMARY KEY,
    name		TEXT NOT NULL,
    
    type_id		BIT(32) NOT NULL UNIQUE,
    weak		BIT(32) NOT NULL,
    resistant	BIT(32) NOT NULL,
    immune		BIT(32) NOT NULL
);
	`)

	if err != nil {
		return err
	}

	_, err = database.Exec(`
CREATE TABLE IF NOT EXISTS public.movepools (
    id			BIGSERIAL PRIMARY KEY,
    monster_id	BIGINT NOT NULL REFERENCES public.base_monsters(monster_id),
    level		SMALLINT NOT NULL,
    move_id		SMALLINT NOT NULL
);
	`)

	if err != nil {
		return err
	}

	_, err = database.Exec(`
CREATE TABLE IF NOT EXISTS public.tamed_monsters (
    id			BIGSERIAL PRIMARY KEY,
    user_id		BIT(64) NOT NULL,
    base_id		BIGINT NOT NULL REFERENCES public.base_monsters(monster_id),
    
    ev_hp		SMALLINT NOT NULL DEFAULT 0,
    ev_attack	SMALLINT NOT NULL DEFAULT 0,
    ev_defense	SMALLINT NOT NULL DEFAULT 0,
    ev_sattack	SMALLINT NOT NULL DEFAULT 0,
    ev_sdefense	SMALLINT NOT NULL DEFAULT 0,
    ev_speed	SMALLINT NOT NULL DEFAULT 0,
    
    iv_hp		SMALLINT NOT NULL,
    iv_attack	SMALLINT NOT NULL,
    iv_defense	SMALLINT NOT NULL,
    iv_sattack	SMALLINT NOT NULL,
    iv_sdefense	SMALLINT NOT NULL,
    iv_speed	SMALLINT NOT NULL,
    
    nature		SMALLINT NOT NULL,
    gender		BIT(1) NULL,
    level		SMALLINT NOT NULL DEFAULT 1,
    pokerus		BOOLEAN NOT NULL DEFAULT FALSE,
    nickname	TEXT NULL
);
	`)

	if err != nil {
		return err
	}

	return nil
}
