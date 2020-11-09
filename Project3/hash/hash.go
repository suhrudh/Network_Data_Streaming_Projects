package hash

// import (
// 	"fmt"
// )

import(
		"project3/customrandom")

//HFunctionCollection , array of random numbers
type HFunctionCollection []int

//CreateHashCollection , creates collection of hash functions
func (arr *HFunctionCollection)CreateHashCollection(size  int) {

	for i:=0; i < size; i++{
		(*arr) = append((*arr),customrandom.GetRandomNumber(0,2147483647))
	}
	// println("hashes :")
	// for i:=0; i < size; i++{
	// 	println((*arr)[i])
	// }
	// println("###")
}

func hashFunction( num int) int{
	return (num*(2654435761))>>32
}

//ActHashN - acting nth hash function 
func (arr *HFunctionCollection)ActHashN(n , flowID, tableSize int) int{
	temp := (*arr)[n]
	return ((hashFunction(flowID)^(temp))%tableSize)
}

//ActHashNWithoutMod - acting nth hash function 
func (arr *HFunctionCollection)ActHashNWithoutMod(n , flowID int) int{
	temp := (*arr)[n]
	return (hashFunction(flowID)^temp)
}