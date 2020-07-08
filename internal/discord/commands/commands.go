package commands

import (
	"github.com/bwmarrin/discordgo"
	"math"
	"strings"
)

const OwnerOnly = math.MaxInt64

var CommandList = make(map[string]Command)

func AddCommand(command Command) {
	CommandList[command.Name()] = command
	for _, v := range command.Aliases() {
		CommandList[v] = command
	}
}

func init() {
	AddCommand(RandomPokeCmd{})
	AddCommand(HelpCmd{})
	AddCommand(PokeTypesCmd{})
	AddCommand(AddPokeCmd{})
}

type Command interface {
	Help() string
	Run(session *discordgo.Session, msg *discordgo.Message, args []string)
	Name() string
	Permissions() (required int)
	Aliases() []string
}

func TryInvoke(s *discordgo.Session, msg *discordgo.Message, command Command) {
	reqPerms := command.Permissions()
	app, err := s.Application("@me")

	if err != nil {
		panic(err)
	}

	if perms, err := s.State.UserChannelPermissions(msg.Author.ID, msg.ChannelID); err != nil {
		panic(err)
	} else if reqPerms&perms == reqPerms || msg.Author.ID == app.Owner.ID {
		go command.Run(s, msg, strings.Fields(msg.Content)[1:])
	} else {
		return
	}
}
