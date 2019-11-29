package moves

import (
	"github.com/frederikthedane/DiscordMon/internal/constants"
	"github.com/frederikthedane/DiscordMon/internal/mechanics"
)

var MoveTackle = mechanics.PokeMove{
	Type:           constants.TypeNormal,
	Category:       constants.Physical,
	PP:             35,
	Power:          40,
	Accuracy:       1.0,
	FlinchChance:   0,
	Priority:       0,
	Contact:        true,
	OwnStatChanges: [6]int{},
	OppStatChanges: [6]int{},
}
