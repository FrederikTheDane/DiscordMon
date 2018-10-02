package monsters

import (
	"github.com/frederikthedane/discordmon"
	"github.com/frederikthedane/discordmon/moves"
	"github.com/frederikthedane/discordmon/constants"
	"math/rand"
	"time"
)

var movepool = discordmon.MovePool{
	0: []discordmon.PokeMove{discordmonMoves.MoveAstonish},
	3: []discordmon.PokeMove{},
}

var base = discordmon.PokeMonBase{
	Gendered:   false,
	ID:         1,
	Types:      [2]constants.PokeType{},
	BaseStats:  [6]int{40, 30, 30, 40, 50, 30},
	EVYield:    [6]int{ 0,  0,  0,  0,  1,  0},
	MovePool:   movepool,
	Name:       ":frog:",
}

func newFrog() *discordmon.PokeMon {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIVS := *new([6]int)
	moves := *new([4]discordmon.PokeMove)
	evs := *new([6]int)
	for i := 0; i < cap(randomIVS); i++ {
		randomIVS[i] = random.Intn(32)
	}

	mon := discordmon.PokeMon{
		Base:       base,
		EVS:        evs,
		IVS:        randomIVS,
		StatStages: [6]int{},
		Stats:      [6]int{},
		Moves:      moves,
		Gender:     random.Intn(2),
		NVStatus:   0,
		VStatus:    0,
		Level:      1,
		Nature:     random.Intn(constants.NatureSerious + 1),
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
	base.New = newFrog
	discordmon.Table[base.ID] = base
}
