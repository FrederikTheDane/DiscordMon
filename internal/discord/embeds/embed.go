package embeds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
	"strconv"
)

func CreateMonsterEmbed(mon *mechanics.PokeMon) discordgo.MessageEmbed {
	stats := mon.CalculateStats()
	fields := new([6]*discordgo.MessageEmbedField)
	for i, _ := range fields {
		fields[i] = &discordgo.MessageEmbedField{
			Name:   constants.StatNames[constants.Stat(i)],
			Value:  strconv.Itoa(stats[i]),
			Inline: false,
		}
	}

	embed := discordgo.MessageEmbed{
		Title:       mon.Nick,
		Description: mon.Base.Name,
		Color:       0xECEC00,
		Fields:      fields[:],
	}
	return embed
}

func CreateBaseMonsterEmbed(mon *mechanics.PokeMonBase) discordgo.MessageEmbed {
	fields := new([6]*discordgo.MessageEmbedField)
	for i, _ := range fields {
		fields[i] = &discordgo.MessageEmbedField{
			Name:   constants.StatNames[constants.Stat(i)],
			Value:  strconv.Itoa(mon.BaseStats[i]),
			Inline: false,
		}
	}

	embed := discordgo.MessageEmbed{
		Title:       mon.Name,
		Description: "Base stats",
		Color:       0x07CA07,
		Fields:      fields[:],
	}

	return embed
}

func CreateTypeEmbed(t *constants.PokeType) discordgo.MessageEmbed {
	weaknesses := make([]constants.PokeType, 0)
	resistances := make([]constants.PokeType, 0)
	immunities := make([]constants.PokeType, 0)

	//Check for strengths / weaknesses
	for _, v := range constants.TypeMap {
		if (v.TypeID & t.WeakDef) > 0 {
			weaknesses = append(weaknesses, v)
		}
		if (v.TypeID & t.StrongDef) > 0 {
			resistances = append(resistances, v)
		}
		if (v.TypeID & t.NoEffect) > 0 {
			immunities = append(immunities, v)
		}
	}

	weakField := discordgo.MessageEmbedField{
		Name:   "Weaknesses",
		Value:  fmt.Sprint(weaknesses),
		Inline: false,
	}

	resistantField := discordgo.MessageEmbedField{
		Name:   "Resistances",
		Value:  fmt.Sprint(resistances),
		Inline: false,
	}

	noEffectField := discordgo.MessageEmbedField{
		Name:   "No effect",
		Value:  fmt.Sprint(immunities),
		Inline: false,
	}

	fields := [...]*discordgo.MessageEmbedField{&weakField, &resistantField, &noEffectField}

	embed := discordgo.MessageEmbed{
		Title:  t.Name,
		Color:  0x3C65D6,
		Fields: fields[:],
	}

	return embed
}
