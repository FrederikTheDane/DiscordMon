package discordmonMoves

import (
	"github.com/frederikthedane/discordmon"
	"github.com/frederikthedane/discordmon/constants"
)

var MoveAbsorb = discordmon.PokeMove{
	Type:           constants.TypeGrass,
	Category:       constants.CategorySpecial,
	PP:             25,
	Power:          20,
	Accuracy:       1.0,
	FlinchChance:   0,
	Priority:       0,
	Contact:        false,
	OwnStatChanges: [6]int{},
	OppStatChanges: [6]int{},
}
