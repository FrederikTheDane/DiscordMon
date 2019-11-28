package moves

import (
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
)

var MoveAstonish = mechanics.PokeMove{
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
