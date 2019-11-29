package monsters

import (
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
	"github.com/frederikthedane/DiscordMon/internal/moves"
	"math/rand"
	"time"
)

var movepool = mechanics.MovePool{
	0: []mechanics.PokeMove{moves.MoveAstonish},
	3: []mechanics.PokeMove{},
}

var base = mechanics.PokeMonBase{
	Gendered:  false,
	ID:        1,
	Types:     [2]constants.PokeType{},
	BaseStats: [6]int{40, 30, 30, 40, 50, 30},
	EVYield:   [6]int{0, 0, 0, 0, 1, 0},
	MovePool:  movepool,
	Name:      ":frog:",
}

func NewFrog() *mechanics.PokeMon {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIVS := *new([6]int)
	pokeMoves := *new([4]mechanics.PokeMove)
	evs := *new([6]int)
	for i := 0; i < cap(randomIVS); i++ {
		randomIVS[i] = random.Intn(32)
	}

	mon := mechanics.PokeMon{
		Base:       base,
		EVS:        evs,
		IVS:        randomIVS,
		StatStages: [6]int{},
		Stats:      [6]int{},
		Moves:      pokeMoves,
		Gender:     random.Intn(2),
		NVStatus:   0,
		VStatus:    0,
		Level:      1,
		Nature:     constants.Nature(random.Intn(int(constants.Serious) + 1)),
		PokeRus:    false,
		Nick:       "Frog",
	}
	mon.Stats = mon.CalculateStats()
	return &mon
}

func init() {
	types := *new([2]constants.PokeType)
	types[0] = constants.TypeWater
	types[1] = constants.TypeGrass
	base.Types = types
	base.New = NewFrog
}
