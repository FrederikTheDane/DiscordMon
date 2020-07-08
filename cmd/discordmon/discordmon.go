package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/discord/commands"
	"github.com/frederikthedane/DiscordMon/internal/discord/database"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
	"github.com/frederikthedane/DiscordMon/internal/monsters"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	token := flag.String("token", "", "Bot token")
	dbUser := flag.String("dbuser", "", "Database user")
	dbPwd := flag.String("dbpwd", "", "Database password")
	dbHost := flag.String("dbhost", "localhost", "Database host network location")
	dbPort := flag.String("dbport", "5432", "Database host port")
	dbName := flag.String("dbname", "discordmon", "Database name")

	flag.Parse()

	if *dbUser == "" || *dbPwd == "" {
		log.Fatalln("Please provide a database username and password")
	}

	if *token == "" {
		log.Fatalln("Please provide a token")
	}

	os.Setenv("DISCORDMON_DB_HOST", *dbHost)
	os.Setenv("DISCORDMON_DB_PORT", *dbPort)
	os.Setenv("DISCORDMON_DB_USER", *dbUser)
	os.Setenv("DISCORDMON_DB_PWD", *dbPwd)
	os.Setenv("DISCORDMON_DB_NAME", *dbName)

	dbErr := database.Connect()
	if dbErr != nil {
		log.Fatalln(dbErr)
	}

	var mons = make([]*mechanics.PokeMon, 2)

	//displayNatures()
	fmt.Printf("\n")
	for k, _ := range mons {
		mons[k] = monsters.NewFrog()
		mons[k].SetLevel(100)
		mons[k].GainEVs([6]int{255, 255, 255, 255, 255, 255})
		mons[k].SetNick("Test subject " + fmt.Sprintf("%v", k))
		//displayPokemon(v, false)
	}

	session, err := discordgo.New("Bot " + *token)
	if err != nil {
		panic(err)
	}

	session.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		fmt.Printf("Logged in as %s#%s\n", e.User.Username, e.User.Discriminator)
		fmt.Printf("Connected to %d guilds\n", len(e.Guilds))
	})

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, "dm!") {
			cmdName := strings.Fields(m.Content[3:])[0]
			if cmd, ok := commands.CommandList[cmdName]; ok {
				commands.TryInvoke(s, m.Message, cmd)
			}
		}
	})

	if err = session.Open(); err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sig

	_ = session.Close()
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
