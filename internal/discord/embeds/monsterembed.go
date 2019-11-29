package embeds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
)

func CreateMonsterEmbed(mon *mechanics.PokeMon) discordgo.MessageEmbed {
	embed := discordgo.MessageEmbed{
		URL:         "",
		Type:        "",
		Title:       mon.Nick,
		Description: mon.Base.Name,
		Timestamp:   "",
		Color:       0xECEC00,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
		Fields:      nil,
	}
	return embed
}
