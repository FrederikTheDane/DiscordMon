package constants

type PokeType struct {
	Name      string
	TypeID    int
	WeakDef   int
	StrongDef int
	NoEffect  int
}

// All Pokemon types, as a mask
const (
	IDNormal = 1 << iota
	IDFire
	IDWater
	IDElectric
	IDGrass
	IDIce
	IDFighting
	IDPoison
	IDGround
	IDFlying
	IDPsychic
	IDBug
	IDRock
	IDGhost
	IDDragon
	IDDark
	IDSteel
	IDFairy
)

var (
	TypeNormal = PokeType{
		Name:      "Normal",
		TypeID:    IDNormal,
		WeakDef:   IDFighting,
		StrongDef: 0,
		NoEffect:  IDGhost,
	}

	TypeFire = PokeType{
		Name:      "Fire",
		TypeID:    IDFire,
		WeakDef:   IDWater | IDGround | IDRock,
		StrongDef: IDFire | IDGrass | IDIce | IDBug | IDSteel | IDFairy,
		NoEffect:  0,
	}

	TypeWater = PokeType{
		Name:      "Water",
		TypeID:    IDWater,
		WeakDef:   IDElectric | IDGrass,
		StrongDef: IDFire | IDWater | IDIce | IDSteel,
		NoEffect:  0}

	TypeElectric = PokeType{
		Name:      "Electric",
		TypeID:    IDElectric,
		WeakDef:   IDGround,
		StrongDef: IDElectric | IDFlying | IDSteel,
		NoEffect:  0}

	TypeGrass = PokeType{
		Name:      "Grass",
		TypeID:    IDGrass,
		WeakDef:   IDFire | IDIce | IDPoison | IDFlying | IDBug,
		StrongDef: IDWater | IDElectric | IDGrass | IDGround,
		NoEffect:  0}

	TypeIce = PokeType{
		Name:      "Ice",
		TypeID:    IDIce,
		WeakDef:   IDFire | IDFighting | IDRock | IDSteel,
		StrongDef: IDIce,
		NoEffect:  0}

	TypeFighting = PokeType{
		Name:      "Fighting",
		TypeID:    IDFighting,
		WeakDef:   IDFlying | IDPsychic | IDFairy,
		StrongDef: IDBug | IDRock | IDDark,
		NoEffect:  0}

	TypePoison = PokeType{
		Name:      "Poison",
		TypeID:    IDPoison,
		WeakDef:   IDGround | IDPsychic,
		StrongDef: IDGrass | IDFighting | IDPoison | IDBug | IDFairy,
		NoEffect:  0}

	TypeGround = PokeType{
		Name:      "Ground",
		TypeID:    IDGround,
		WeakDef:   IDWater | IDGrass | IDIce,
		StrongDef: IDPoison | IDGround,
		NoEffect:  IDElectric}

	TypeFlying = PokeType{
		Name:      "Flying",
		TypeID:    IDFlying,
		WeakDef:   IDElectric | IDIce | IDRock,
		StrongDef: IDGrass | IDFighting | IDBug,
		NoEffect:  IDGround}

	TypePsychic = PokeType{
		Name:      "Psychic",
		TypeID:    IDPsychic,
		WeakDef:   IDBug | IDGhost | IDDark,
		StrongDef: IDFighting | IDPsychic,
		NoEffect:  0}

	TypeBug = PokeType{
		Name:      "Bug",
		TypeID:    IDBug,
		WeakDef:   IDFire | IDFlying | IDRock,
		StrongDef: IDGrass | IDFighting | IDGround,
		NoEffect:  0}

	TypeRock = PokeType{
		Name:      "Rock",
		TypeID:    IDRock,
		WeakDef:   IDWater | IDGrass | IDFighting | IDGround | IDSteel,
		StrongDef: IDNormal | IDFire | IDPoison | IDFlying,
		NoEffect:  0}

	TypeGhost = PokeType{
		Name:      "Ghost",
		TypeID:    IDGhost,
		WeakDef:   IDGhost | IDDark,
		StrongDef: IDPoison | IDBug,
		NoEffect:  IDFighting | IDNormal}

	TypeDragon = PokeType{
		Name:      "Dragon",
		TypeID:    IDDragon,
		WeakDef:   IDIce | IDDragon | IDFairy,
		StrongDef: IDFire | IDWater | IDElectric | IDGrass,
		NoEffect:  0}

	TypeDark = PokeType{
		Name:      "Dark",
		TypeID:    IDDark,
		WeakDef:   IDFighting | IDBug | IDFairy,
		StrongDef: IDGhost | IDDark,
		NoEffect:  IDPsychic}

	TypeSteel = PokeType{
		Name:      "Steel",
		TypeID:    IDSteel,
		WeakDef:   IDFire | IDFighting | IDGround,
		StrongDef: IDNormal | IDGrass | IDIce | IDFlying | IDPsychic | IDBug | IDRock | IDDragon | IDSteel | IDFairy,
		NoEffect:  IDPoison}

	TypeFairy = PokeType{
		Name:      "Fairy",
		TypeID:    IDFairy,
		WeakDef:   IDPoison | IDSteel,
		StrongDef: IDFighting | IDBug | IDDark,
		NoEffect:  IDDragon}
)

var (
	TypeMap map[int]PokeType
)

func init() {
	TypeMap = make(map[int]PokeType)

	TypeMap[IDNormal] = TypeNormal
	TypeMap[IDFire] = TypeFire
	TypeMap[IDWater] = TypeWater
	TypeMap[IDElectric] = TypeElectric
	TypeMap[IDGrass] = TypeGrass
	TypeMap[IDIce] = TypeIce
	TypeMap[IDFighting] = TypeFighting
	TypeMap[IDPoison] = TypePoison
	TypeMap[IDGround] = TypeGround
	TypeMap[IDFlying] = TypeFlying
	TypeMap[IDPsychic] = TypePsychic
	TypeMap[IDBug] = TypeBug
	TypeMap[IDRock] = TypeRock
	TypeMap[IDGhost] = TypeGhost
	TypeMap[IDDragon] = TypeDragon
	TypeMap[IDDark] = TypeDark
	TypeMap[IDSteel] = TypeSteel
	TypeMap[IDFairy] = TypeFairy
}
