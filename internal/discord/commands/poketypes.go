package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/discord/embeds"
	"log"
)

type PokeTypesCmd struct {
}

func (p PokeTypesCmd) Help() string {
	return "Show type matchups for all types (Sends in DM to avoid clutter)"
}

func (p PokeTypesCmd) Run(session *discordgo.Session, msg *discordgo.Message, args []string) {
	for _, v := range constants.TypeMap {
		embed := embeds.CreateTypeEmbed(&v)
		if ch, err := session.UserChannelCreate(msg.Author.ID); err != nil {
			log.Println(err)
		} else {
			if _, err = session.ChannelMessageSendEmbed(ch.ID, &embed); err != nil {
				log.Println(err)
			}
		}
	}
}

func (p PokeTypesCmd) Name() string {
	return "types"
}

func (p PokeTypesCmd) Permissions() (required int) {
	return 0
}

func (p PokeTypesCmd) Aliases() []string {
	return make([]string, 0)
}
