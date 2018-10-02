package discordmonMoves

import (
	"github.com/frederikthedane/discordmon"
	"github.com/frederikthedane/discordmon/constants"
)

var MoveAstonish = discordmon.PokeMove{
	Type:           constants.TypeGhost,
	Category:       constants.CategoryPhysical,
	PP:             15,
	Power:          30,
	Accuracy:       1.0,
	FlinchChance:   0.3,
	Priority:       0,
	Contact:        true,
	OwnStatChanges: [6]int{},
	OppStatChanges: [6]int{},
}
