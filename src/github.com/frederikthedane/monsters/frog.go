package discordmonMonsters

import (
	"math/rand"
	"time"
	"github.com/frederikthedane/discordmon"
	"github.com/frederikthedane/moves"
)

var base = discordmon.PokeMonBase{
	Gendered:   false,
	ID:         1,
	Types:      [2]discordmon.PokeType{},
	BaseStats:  [6]int{40, 30, 30, 40, 50, 30},
	EVYield:    [6]int{ 0,  0,  0,  0,  1,  0},
	MovePool:   discordmon.MovePool{},
	EggGroup:   discordmon.EggGroup{},
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
	moves[0] = discordmonMoves.TackleMove


	mon := discordmon.PokeMon{
		Base:     base,
		EVS:      evs,
		IVS:      randomIVS,
		Stats:    [6]int{},
		Moves:    moves,
		Gender:   random.Intn(2),
		NVStatus: 0,
		VStatus:  0,
		Level:    1,
		Nature:   random.Intn(discordmon.NatureSerious + 1),
		PokeRus:  false,
	}
	mon.Stats = mon.CalculateStats()
	return &mon
}

func init() {
	types := *new([2]discordmon.PokeType)
	types[0] = discordmon.TypeWater
	types[1] = discordmon.TypeGrass
	base.Types = types
	base.New = newFrog
	discordmon.Table[base.ID] = base
}