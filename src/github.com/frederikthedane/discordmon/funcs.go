package discordmon

var Table map[int]PokeMonBase

func init() {
	Table = map[int]PokeMonBase{}
}

func sumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

//Get a new PokeMon with the given numeric ID
func NewFromID(ID int) *PokeMon {
	return Table[ID].New()
}

//TODO: Finish this func and test it
func DoTurn(mon1 *PokeMon, move1 PokeMove, mon2 *PokeMon, move2 PokeMove) {

}