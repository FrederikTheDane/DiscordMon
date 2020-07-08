package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/frederikthedane/DiscordMon/internal/discord/embeds"
	"github.com/frederikthedane/DiscordMon/internal/monsters"
	"log"
)

type RandomPokeCmd struct {
}

func (r RandomPokeCmd) Help() string {
	return "Shows a random pokemon (mostly for testing)"
}

func (r RandomPokeCmd) Run(session *discordgo.Session, msg *discordgo.Message, args []string) {
	embed := embeds.CreateMonsterEmbed(monsters.NewFrog())
	if _, err := session.ChannelMessageSendEmbed(msg.ChannelID, &embed); err != nil {
		log.Println(err)
	}
}

func (r RandomPokeCmd) Name() string {
	return "randpoke"
}

func (r RandomPokeCmd) Permissions() (required int) {
	return 0
}

func (r RandomPokeCmd) Aliases() []string {
	return make([]string, 0)
}
