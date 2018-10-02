package discordmonMoves

import (
	"github.com/frederikthedane/discordmon"
	"github.com/frederikthedane/discordmon/constants"
)


var MoveTackle = discordmon.PokeMove{
	Type:           constants.TypeNormal,
	Category:       constants.CategoryPhysical,
	PP:             35,
	Power:          40,
	Accuracy:       1.0,
	FlinchChance:   0,
	Priority:       0,
	Contact:        true,
	OwnStatChanges: [6]int{},
	OppStatChanges: [6]int{},
}