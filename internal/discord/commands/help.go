package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type HelpCmd struct {
}

func (h HelpCmd) Help() string {
	return "Display help info for a given command"
}

func (h HelpCmd) Run(session *discordgo.Session, msg *discordgo.Message, args []string) {
	var help string
	if len(args) == 0 {
		help = h.Help()
	} else {
		if cmd, ok := CommandList[args[0]]; ok {
			help = cmd.Help()
		} else {
			help = fmt.Sprintf("Command not found: %s", args[0])
		}
	}
	if _, err := session.ChannelMessageSend(msg.ChannelID, help); err != nil {
		log.Println(err)
	}
}

func (h HelpCmd) Name() string {
	return "help"
}

func (h HelpCmd) Permissions() (required int) {
	return 0
}

func (h HelpCmd) Aliases() []string {
	return make([]string, 0)
}
