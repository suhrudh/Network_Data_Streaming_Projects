package hash

// import (
// 	"fmt"
// )

import(
		"project2/customrandom")

//HFunctionCollection , array of random numbers
type HFunctionCollection []int

//CreateHashCollection , creates collection of hash functions
func (arr *HFunctionCollection)CreateHashCollection(size  int) {

	for i:=0; i < size; i++{
		(*arr) = append((*arr),customrandom.GetRandomNumber(0,2147483647))
	}

	// for i:=0; i < size; i++{
	// 	fmt.Println((*arr)[i])
	// }
}

func hashFunction( num int) int{
	return num*(2654435761)
}

//ActHashN - acting nth hash function 
func (arr *HFunctionCollection)ActHashN(n , flowID, tableSize int) int{
	temp := (*arr)[n]
	return ((hashFunction(flowID)^(temp))%tableSize)
}