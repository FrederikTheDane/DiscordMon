package discordmon

var Table map[int]PokeMonBase

func init() {
	Table = map[int]PokeMonBase{}
}

//Sums ints in an array
func sumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

//Sums the values in two arrays
func sumArrays(arr1, arr2 []int) []int {
	if len(arr1) != len(arr2) {
		panic("Panic! The array lengths do not match!")
	}
	sumArr := make([]int, len(arr1))
	for k := range sumArr {
		sumArr[k] = arr1[k] + arr2[k]
	}
	return sumArr
}

//Gets key of the largest int in an array
func getLargestKey(ints []int) int {
	largest := 0
	for k, v := range ints {
		if v > ints[largest] {
			largest = k
		}
	}
	return largest
}

//Get a new PokeMon with the given numeric ID
func NewFromID(ID int) *PokeMon {
	return Table[ID].New()
}

//TODO: Make this func flexible to allow moves with special effects such as absorb, super fang and so on.
//Maybe by using func parameters?
func DoTurn(mon1 *PokeMon, move1 *PokeMove, mon2 *PokeMon, move2 *PokeMove) {

}