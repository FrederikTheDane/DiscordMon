package commands

import (
	"github.com/bwmarrin/discordgo"
)

type ShowPokeCmd struct {
}

func (c ShowPokeCmd) Aliases() []string {
	aliases := make([]string, 0)
	aliases = append(aliases, "sp")
	return aliases
}

func (c ShowPokeCmd) Help() string {
	return "Use this command to show a pokemon"
}

func (c ShowPokeCmd) Run(session *discordgo.Session, msg *discordgo.Message, args []string) {
	if len(args) > 0 {

	} else {

	}
}

func (c ShowPokeCmd) Name() string {
	return "showpoke"
}

func (c ShowPokeCmd) Permissions() (required int) {
	return 0
}
