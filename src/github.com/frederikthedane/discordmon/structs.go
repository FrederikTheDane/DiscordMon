package discordmon

import (
	"math"
	)


//Stuff that cannot change throughout the PokeMons lifecycle
type PokeMonBase struct {
	Gendered bool
	ID, CatchRate, HatchSteps int
	Types [2]PokeType
	BaseStats [6]int
	EVYield [6]int
	MovePool MovePool
	EggGroup EggGroup
	Name string
	Evolutions []PokeMon
	New func() *PokeMon
}

//Things that may change for a pokemon over time
type PokeMon struct {
	Base PokeMonBase
	EVS [6]int
	IVS [6]int
	Stats [6]int
	Moves [4]PokeMove
	Gender, NVStatus, VStatus, Level, Nature int
	PokeRus bool
}

func (mon *PokeMon) SetLevel(lvl int) {
	mon.Level = lvl
	mon.Stats = mon.CalculateStats()
}

//TODO: Fix this so it will round down if the new EV's exceed maximum
func (mon *PokeMon) GainEVs(evs [6]int) {
	tempEVS := mon.EVS

	exceedsMaxEVS := func(newEVs []int) bool {
		if sumInts(newEVs) > 510 {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < len(evs); i++ {
		tempEVS[i] += evs[i]
		nextTry := tempEVS
		nextTry[i] -= 1
		finalTry := nextTry
		finalTry[i] -= 1
		if !exceedsMaxEVS(tempEVS[:]) && tempEVS[i] + evs[i] <= 252{
			mon.EVS[i] += evs[i]
		} else if nextTry[i] + evs[i] <= 252 {
			mon.EVS[i] += nextTry[i]
		} else if finalTry[i] + evs[i] <= 252 {
			mon.EVS[i] += finalTry[i]
		}
	}

	mon.Stats = mon.CalculateStats()
}

//Calculates a PokeMons stats (including hp) from lvl, iv's ev's and base stats
func (mon *PokeMon) CalculateStats() [6]int {
	stats := *new([6]int)
	stats[0] = mon.CalculateHP()
	lvl := mon.Level
	for i := 1; i < 6; i++ {
		bsStat := mon.Base.BaseStats[i]
		ivStat := mon.IVS[i]
		evStat := mon.EVS[i]
		natureMod := float64(Natures[mon.Nature][i])
		stats[i] = int(math.Floor(math.Floor(float64((2 * bsStat + ivStat + evStat) * lvl / 100 + 5) * natureMod)))
	}
	return stats
}

//Same as above, but for only HP as the formula is a bit different
func (mon *PokeMon) CalculateHP() int {
	bsHP := mon.Base.BaseStats[0]
	ivHP := mon.IVS[0]
	evHP := mon.EVS[0]
	lvl := mon.Level
	stat := int(math.Floor(float64((2 * bsHP + ivHP + evHP) * lvl / 100 + lvl + 10)))
	return stat
}

//TODO: Test this func. It calculates the effectiveness compared to the RECEIVING pokemon
func (mon *PokeMon) CalculateDamageMultiplier(pokeType PokeType) float64 {
	types := mon.Base.Types
	multiplier := 1.0
	for i := 0; i < len(types); i++ {
		if types[i].WeakDef & pokeType.TypeID > 0 {
			multiplier *= 2
		} else if types[i].StrongDef & pokeType.TypeID > 0 {
			multiplier /= 2
		} else if types[i].NoEffect & pokeType.TypeID > 0 {
			return 0
		}
	}
	return multiplier
}


type PokeMove struct {
	Type PokeType
	Category int
	PP int
	Power int
	Accuracy float64
	Priority int
	Contact bool
}

type PokeType struct {
	Name string
	TypeID int
	WeakDef int
	StrongDef int
	NoEffect int
}

type MovePool struct {
	Moves map[int][]PokeMove
}

type EggGroup struct {

}
