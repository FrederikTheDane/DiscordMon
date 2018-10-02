package main

import (
	"fmt"
	"github.com/frederikthedane/discordmon"
	"github.com/frederikthedane/discordmon/constants"
	_ "github.com/frederikthedane/discordmon/monsters"
	"github.com/frederikthedane/discordmon/moves"
)

func main() {
	var monsters = make([]*discordmon.PokeMon, 2)

	displayNatures()
	fmt.Printf("\n")
	for k, v := range monsters{
		v = discordmon.NewFromID(1)
		v.SetLevel(100)
		v.GainEVs([6]int{255, 255, 255, 255, 255, 255})
		v.SetNick("Test subject " + fmt.Sprintf("%v", k))
		displayPokemon(v, false)
	}

	discordmon.DoTurn(monsters[0], &discordmonMoves.MoveTackle, monsters[1], &discordmonMoves.MoveTackle)
}

func displayNatures() {
	for i := 0; i <= constants.NatureSerious; i++ {
		fmt.Printf("%10s %4v\n", constants.NatureNames[i], constants.Natures[i])
	}
}

func displayPokemon(mon *discordmon.PokeMon, detailed bool) {
	fmt.Printf("%16s %s\n", "Pokemon name:", mon.Base.Name)
	fmt.Printf("%16s %s\n", "Pokemon nick:", mon.Nick)
	fmt.Printf("%16s %d\n", "Level:", mon.Level)

	if detailed {
		fmt.Printf("%16s %3v\n","Base stats:", mon.Base.BaseStats)
		fmt.Printf("%16s %3v\n", "Pokemon IV's:", mon.IVS)
		fmt.Printf("%16s %3v\n", "Pokemon EV's:", mon.EVS)

	}

	fmt.Printf("%16s %3v\n", "Pokemon stats:", mon.Stats)
	fmt.Printf("%16s %s\n", "Nature:", constants.NatureNames[mon.Nature])
	fmt.Printf("%16s %s, %s\n\n", "Types:", mon.Base.Types[0].Name, mon.Base.Types[1].Name)

	if detailed {
		fmt.Printf("%16s\n", "Dmg received:")

		for _, v := range constants.TypeMap {
			fmt.Printf("%15s: %.2f\n", v.Name, mon.CalculateDamageMultiplier(v))
		}
	}
}
