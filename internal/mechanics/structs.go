package mechanics

import (
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"math"
)


//Stuff that cannot change throughout the PokeMons lifecycle
type PokeMonBase struct {
	Gendered bool
	ID, CatchRate, HatchSteps int
	Types [2]constants.PokeType
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
	StatStages [6]int
	Stats [6]int
	BattleStats [6]int
	Moves [4]PokeMove
	Gender, NVStatus, VStatus, Level, Nature int
	PokeRus bool
	Nick string
}

//Gives the pokemon a nickname
func (mon *PokeMon) SetNick(nick string) {
	mon.Nick = nick
}

//Sets the pokemon to a given level, and calculates its new stats
func (mon *PokeMon) SetLevel(lvl int) {
	if lvl > 100 {
		mon.Level = 100
	} else {
		mon.Level = lvl
	}
	mon.Stats = mon.CalculateStats()
}

//Rounds down if the new EV's exceed maximum. 252 for each stat max, 510 for all
func (mon *PokeMon) GainEVs(evs [6]int) {

	exceedsMaxEVS := func(newEVs []int) bool {
		for _, v := range newEVs {
			if v > 252 {
				return true
			}
		}
		return sumInts(newEVs) > 510
	}

	var addEVs func(toAdd []int) []int
	addEVs = func(toAdd []int) []int {
		adder := toAdd
		tempEVs := toAdd
		if exceedsMaxEVS(sumArrays(mon.EVS[:], tempEVs[:])) {
			tempEVs[getLargestKey(toAdd[:])] -= 1
			adder = addEVs(tempEVs)

		}
		return adder
	}
	copy(mon.EVS[:], sumArrays(addEVs(evs[:]), mon.EVS[:]))
	mon.Stats = mon.CalculateStats()
}

//Calculates a PokeMons stats (including hp) from lvl, iv's ev's and base stats
func (mon *PokeMon) CalculateStats() [6]int {
	stats := *new([6]int)
	stats[constants.StatHP] = mon.CalculateHP()
	lvl := mon.Level
	for i := 1; i < 6; i++ {
		bsStat := mon.Base.BaseStats[i]
		ivStat := mon.IVS[i]
		evStat := mon.EVS[i]
		natureMod := float64(constants.Natures[mon.Nature][i])
		stats[i] = int(math.Floor(float64((2 * bsStat + ivStat + evStat) * lvl / 100 + 5) * natureMod))
	}
	return stats
}

//Same as above, but only for HP, as the formula is a bit different
func (mon *PokeMon) CalculateHP() int {
	bsHP := mon.Base.BaseStats[constants.StatHP]
	ivHP := mon.IVS[constants.StatHP]
	evHP := mon.EVS[constants.StatHP]
	lvl := mon.Level
	stat := int(math.Floor(float64((2 * bsHP + ivHP + evHP) * lvl / 100 + lvl + 10)))
	return stat
}

//Calculates the effectiveness compared to the RECEIVING pokemon
func (mon *PokeMon) CalculateDamageMultiplier(pokeType constants.PokeType) float64 {
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
	Type constants.PokeType
	Category int
	PP int
	Power int
	Accuracy float64
	FlinchChance float64
	Priority int
	Contact bool
	OwnStatChanges [6]int
	OppStatChanges [6]int
	UseMove func(state PokeBattleState) PokeBattleState
}


//TODO: Fill this struct with relevant entries for an item
type PokeItem struct {
	Name string
	ItemEffect func(state PokeBattleState) PokeBattleState
}

type PokeHeldItem struct {
	Name string
	Condition PokeBattleState
	ItemBattleEffect func(state PokeBattleState) PokeBattleState
}

//TODO: Fill this struct with relevant entries for an ability
type PokeAbility struct {
	Name string
	AbilityBattleEffect func(state PokeBattleState) PokeBattleState

}

//TODO: Fill this struct with relevant information of a state of battle
type PokeBattleState struct {
	Mon1, Mon2 PokeMon
	Turn int
	Weather struct{
		Turns int

	}
}

type MovePool map[int][]PokeMove

type EggGroup int
