package moves

import (
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
)

var MoveAbsorb = mechanics.PokeMove{
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
