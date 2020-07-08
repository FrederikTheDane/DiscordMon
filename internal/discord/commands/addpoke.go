package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/frederikthedane/DiscordMon/internal/discord/database"
	"github.com/frederikthedane/DiscordMon/internal/discord/embeds"
	"strconv"
)

type AddPokeCmd struct {
}

func (a AddPokeCmd) Help() string {
	return "Usage: id: integer, name: string, gendered: bool, types: integer, hp: int, attack: int, defense: integer, special attack: integer, special defense: integer, speed: integer"
}

func (a AddPokeCmd) Run(session *discordgo.Session, msg *discordgo.Message, args []string) {
	success := true

	if len(args) != 10 {
		_, err := session.ChannelMessageSend(msg.ChannelID, "Please provide exactly 10 args. (Check help message for exact parameters)")
		if err != nil {
			panic(err)
		}
		return
	}
	id, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse id as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	name := args[1]

	if len(name) > 256 {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Name must be at most 256 characters")
		if err != nil {
			panic(err)
		}
	}

	gendered, err := strconv.ParseBool(args[2])

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse gendered as boolean: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	types, err := strconv.ParseInt(args[3], 10, 32)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse types as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	hp, err := strconv.ParseInt(args[4], 10, 16)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse hp as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	att, err := strconv.ParseInt(args[5], 10, 16)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse attack as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	def, err := strconv.ParseInt(args[6], 10, 16)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse defense as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	satt, err := strconv.ParseInt(args[7], 10, 16)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse special attack as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	sdef, err := strconv.ParseInt(args[8], 10, 16)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse special defense as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	speed, err := strconv.ParseInt(args[9], 10, 16)

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Could not parse speed as integer: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	mon, err := database.AddNewPokemon(int(id), name, gendered, int(types), int(hp), int(att), int(def), int(satt), int(sdef), int(speed))

	if err != nil {
		success = false
		_, err = session.ChannelMessageSend(msg.ChannelID, "Error inserting into database: "+err.Error())
		if err != nil {
			panic(err)
		}
	}

	if !success {
		_, err = session.ChannelMessageSend(msg.ChannelID, "Error(s) creating pokemon")
		if err != nil {
			panic(err)
		}
		return
	}

	embed := embeds.CreateBaseMonsterEmbed(mon)

	_, err = session.ChannelMessageSend(msg.ChannelID, "Pokemon created!:")
	if err != nil {
		panic(err)
	}

	_, err = session.ChannelMessageSendEmbed(msg.ChannelID, &embed)
	if err != nil {
		panic(err)
	}
}

func (a AddPokeCmd) Name() string {
	return "addpoke"
}

func (a AddPokeCmd) Permissions() (required int) {
	return 0
}

func (a AddPokeCmd) Aliases() []string {
	return make([]string, 0)
}
