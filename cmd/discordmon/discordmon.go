package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
	"github.com/frederikthedane/DiscordMon/internal/monsters"
	"log"
	"os"
)

func main() {
	token := *flag.String("token", "", "Bot token")

	if token == "" {
		log.Println("Please provide a token")
		os.Exit(1)
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	session.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		fmt.Printf("Logged in as %s#%s\n", e.User.Username, e.User.Discriminator)
		fmt.Printf("Connected to %d guilds\n", len(e.Guilds))
	})

	var mons = make([]*mechanics.PokeMon, 2)

	displayNatures()
	fmt.Printf("\n")
	for k, v := range mons {
		v = monsters.NewFrog()
		v.SetLevel(100)
		v.GainEVs([6]int{255, 255, 255, 255, 255, 255})
		v.SetNick("Test subject " + fmt.Sprintf("%v", k))
		displayPokemon(v, false)
	}
}

func displayNatures() {
	for i := constants.Hardy; i <= constants.Serious; i++ {
		fmt.Printf("%10s %4v\n", constants.NatureNames[i], constants.Natures[i])
	}
}

func displayPokemon(mon *mechanics.PokeMon, detailed bool) {
	fmt.Printf("%16s %s\n", "Pokemon name:", mon.Base.Name)
	fmt.Printf("%16s %s\n", "Pokemon nick:", mon.Nick)
	fmt.Printf("%16s %d\n", "Level:", mon.Level)

	if detailed {
		fmt.Printf("%16s %3v\n", "Base stats:", mon.Base.BaseStats)
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
