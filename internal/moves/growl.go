package moves

import (
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
)

var MoveGrowl = mechanics.PokeMove{
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
