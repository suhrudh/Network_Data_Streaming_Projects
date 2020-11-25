package set

import (
	"project4/customrandom"
)

//CreateListRandomNumbersUniqueTo : 
func CreateListRandomNumbersUniqueTo(fullSet *map[int] bool, size int) *[]int{
	var arr []int
	curSize := len(*fullSet)
	for i:=0; i < size; i++{
		temp := customrandom.GetRandomNumber(0,2147483647)
		(*fullSet)[temp] = true
		for curSize + i + 1 != len(*fullSet){
			temp := customrandom.GetRandomNumber(0,2147483647)
			(*fullSet)[temp] = true
		}
		arr = append(arr,temp)
	}
	return &arr
}

	