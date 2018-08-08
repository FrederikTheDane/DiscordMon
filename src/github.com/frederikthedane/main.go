package main

import (
	"github.com/frederikthedane/discordmon"
	_ "github.com/frederikthedane/monsters"
	"fmt"
)

func main() {
	var monster *discordmon.PokeMon
	monster = discordmon.NewFromID(1)
	monster.SetLevel(100)
	monster.GainEVs([6]int{255, 255, 255, 255, 255, 255})
	displayNatures()
	displayPokemon(monster, true)
}

func displayNatures() {
	for i := 0; i <= discordmon.NatureSerious; i++ {
		fmt.Printf("%10s %4v\n", discordmon.NatureNames[i], discordmon.Natures[i])
	}
}

func displayPokemon(mon *discordmon.PokeMon, detailed bool) {
	fmt.Printf("\n")
	fmt.Printf("%16s %s\n", "Pokemon name:", mon.Base.Name)
	fmt.Printf("%16s %d\n", "Level:", mon.Level)

	if detailed {
		fmt.Printf("%16s %3v\n","Base stats:", mon.Base.BaseStats)
		fmt.Printf("%16s %3v\n", "Pokemon IV's:", mon.IVS)
		fmt.Printf("%16s %3v\n", "Pokemon EV's:", mon.EVS)

	}

	fmt.Printf("%16s %3v\n", "Pokemon stats:", mon.Stats)
	fmt.Printf("%16s %s\n", "Nature:", discordmon.NatureNames[mon.Nature])
	fmt.Printf("%16s %s, %s\n\n", "Types:", mon.Base.Types[0].Name, mon.Base.Types[1].Name)

	if detailed {
		fmt.Printf("%16s\n", "Dmg received:")

		for _, v := range discordmon.TypeMap {
			fmt.Printf("%15s: %.2f\n", v.Name, mon.CalculateDamageMultiplier(v))
		}
	}
}