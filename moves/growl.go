package discordmonMoves

import (
	"github.com/frederikthedane/discordmon"
	"github.com/frederikthedane/discordmon/constants"
)

var MoveGrowl = discordmon.PokeMove{
	Type:           constants.TypeNormal,
	Category:       constants.CategoryStatus,
	PP:             40,
	Power:          0,
	Accuracy:       1.0,
	FlinchChance:   0,
	Priority:       0,
	Contact:        false,
	OwnStatChanges: [6]int{},
	OppStatChanges: [6]int{0, -1, 0, 0, 0, 0},
}
