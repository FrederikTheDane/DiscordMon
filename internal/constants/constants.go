package constants

// Move category consts
type MoveCategory int

const (
	Physical MoveCategory = iota
	Special
	Status
)

// Non-volatile status conditions
type NVStat int

const (
	Burned NVStat = iota
	Frozen
	Paralyzed
	Poisoned
	Badlypoisoned
	Sleeping
)

// Volatile status conditions, as a mask
type VStat int

const (
	Bound VStat = 1 << iota
	CantEscape
	Confused
	Cursed
	Embargoed
	Encored
	Flinched
	HealBlocked
	Identified
	Infatuated
	LeechSeeded
	Nightmare
	PerishSong
	Spooked
	Taunted
	Telekinesis
	Tormented
)

//Stat position
type Stat int

const (
	HP Stat = iota
	Attack
	Defense
	SAttack
	SDefense
	Speed
)

// Natures
type Nature int

const (
	Hardy Nature = iota
	Lonely
	Adamant
	Naughty
	Brave
	Bold
	Docile
	Impish
	Lax
	Relaxed
	Modest
	Mild
	Bashful
	Rash
	Quiet
	Calm
	Gentle
	Careful
	Quirky
	Sassy
	Timid
	Hasty
	Jolly
	Naive
	Serious
)

var (
	StatNames   map[Stat]string
	Natures     map[Nature][]float32
	NatureNames map[Nature]string
)

func init() {
	//Doesn't need to be in it's own block, but improves readability
	{
		StatNames = make(map[Stat]string)

		StatNames[HP] = "HP"
		StatNames[Attack] = "Attack"
		StatNames[Defense] = "Defense"
		StatNames[SAttack] = "Special Attack"
		StatNames[SDefense] = "Special Defense"
		StatNames[Speed] = "Speed"
	}

	//Doesn't need to be in it's own block, but improves readability
	{
		NatureNames = make(map[Nature]string)

		NatureNames[Hardy] = "Hardy"
		NatureNames[Lonely] = "Lonely"
		NatureNames[Adamant] = "Adamant"
		NatureNames[Naughty] = "Naughty"
		NatureNames[Brave] = "Brave"
		NatureNames[Bold] = "Bold"
		NatureNames[Docile] = "Docile"
		NatureNames[Impish] = "Impish"
		NatureNames[Lax] = "Lax"
		NatureNames[Relaxed] = "Relaxed"
		NatureNames[Modest] = "Modest"
		NatureNames[Mild] = "Mild"
		NatureNames[Bashful] = "Bashful"
		NatureNames[Rash] = "Rash"
		NatureNames[Quiet] = "Quiet"
		NatureNames[Calm] = "Calm"
		NatureNames[Gentle] = "Gentle"
		NatureNames[Careful] = "Careful"
		NatureNames[Quirky] = "Quirky"
		NatureNames[Sassy] = "Sassy"
		NatureNames[Timid] = "Timid"
		NatureNames[Hasty] = "Hasty"
		NatureNames[Jolly] = "Jolly"
		NatureNames[Naive] = "Naive"
		NatureNames[Serious] = "Serious"
	}

	//Give all natures their corresponding stat modifiers
	//IIRC it loops over the natures and changes the buffed and crippled stats
	Natures = make(map[Nature][]float32)
	var i Nature
	m := 1
	n := 1
	for i = 0; i <= Serious; i++ {
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
