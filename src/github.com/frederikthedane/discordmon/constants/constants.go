package constants

// Move category consts
const (
	CategoryPhysical = iota
	CategorySpecial
	CategoryStatus
)

// Non-volatile status conditions
const (
	NVStatusBurned = iota
	NVStatusFrozen
	NVStatusParalyzed
	NVStatusPoisoned
	NVStatusBadlyPoisoned
	NVStatusSleeping
)

// Volatile status conditions, as a mask
const (
	VStatusBound = 1 << iota
	VStatusCantEscape
	VStatusConfused
	VStatusCursed
	VStatusEmbargoed
	VStatusEncored
	VStatusFlinched
	VStatusHealBlocked
	VStatusIdentified
	VStatusInfatuated
	VStatusLeechSeeded
	VStatusNightmare
	VStatusPerishSong
	VStatusSpooked
	VStatusTaunted
	VStatusTelekinesis
	VStatusTormented
)

//Stat position
const (
	StatHP = iota
	StatAttack
	StatDefense
	StatSAttack
	StatSDefense
	StatSpeed
)

// Natures
const (
	NatureHardy = iota
	NatureLonely
	NatureAdamant
	NatureNaughty
	NatureBrave
	NatureBold
	NatureDocile
	NatureImpish
	NatureLax
	NatureRelaxed
	NatureModest
	NatureMild
	NatureBashful
	NatureRash
	NatureQuiet
	NatureCalm
	NatureGentle
	NatureCareful
	NatureQuirky
	NatureSassy
	NatureTimid
	NatureHasty
	NatureJolly
	NatureNaive
	NatureSerious
)

var (
	StatNames map[int]string
	Natures map[int][]float32
	NatureNames map[int]string
)

func init() {
	//Doesn't need to be in it's own block, but improves readability
	{
		StatNames = make(map[int]string)

		StatNames[StatHP]			= "HP"
		StatNames[StatAttack]		= "Attack"
		StatNames[StatDefense]		= "Defense"
		StatNames[StatSAttack]		= "Special Attack"
		StatNames[StatSDefense]		= "Special Defense"
		StatNames[StatSpeed]		= "Speed"
	}

	//Doesn't need to be in it's own block, but improves readability
	{
		NatureNames = make(map[int]string)

		NatureNames[NatureHardy]	= "Hardy"
		NatureNames[NatureLonely]	= "Lonely"
		NatureNames[NatureAdamant]	= "Adamant"
		NatureNames[NatureNaughty]	= "Naughty"
		NatureNames[NatureBrave]	= "Brave"
		NatureNames[NatureBold]		= "Bold"
		NatureNames[NatureDocile]	= "Docile"
		NatureNames[NatureImpish]	= "Impish"
		NatureNames[NatureLax]		= "Lax"
		NatureNames[NatureRelaxed]	= "Relaxed"
		NatureNames[NatureModest]	= "Modest"
		NatureNames[NatureMild]		= "Mild"
		NatureNames[NatureBashful]	= "Bashful"
		NatureNames[NatureRash]		= "Rash"
		NatureNames[NatureQuiet]	= "Quiet"
		NatureNames[NatureCalm]		= "Calm"
		NatureNames[NatureGentle]	= "Gentle"
		NatureNames[NatureCareful]	= "Careful"
		NatureNames[NatureQuirky]	= "Quirky"
		NatureNames[NatureSassy]	= "Sassy"
		NatureNames[NatureTimid]	= "Timid"
		NatureNames[NatureHasty]	= "Hasty"
		NatureNames[NatureJolly]	= "Jolly"
		NatureNames[NatureNaive]	= "Naive"
		NatureNames[NatureSerious]	= "Serious"
	}

	Natures = make(map[int][]float32)
	var i int
	m := 1
	n := 1
	for i = 0; i <= NatureSerious; i++ {
		numArray := make([]float32, 6)
		for j := 0; j < 6; j++ {
			numArray[j] = 1.0
		}
		Natures[i] = numArray
		Natures[i][m] += 0.1
		Natures[i][n] -= 0.1
		n++
		if n > 5 {
			n = 1
			m++
		}
		if m > 5 {
			m = 1
		}
	}
}