package discordmonMoves

import "github.com/frederikthedane/discordmon"

var TackleMove = discordmon.PokeMove{
	Type:     discordmon.TypeNormal,
	Category: discordmon.CategoryPhysical,
	PP:       35,
	Power:    40,
	Accuracy: 1.0,
	Priority: 0,
	Contact:  true,
}