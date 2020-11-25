package hash

import (
	//"fmt"
	"math"
)

import(
		"project4/customrandom")

//HFunctionCollection , array of random numbers
type HFunctionCollection []int

//CreateHashCollection , creates collection of hash functions
func (hFunc *HFunctionCollection)CreateHashCollection(size  int, fullSet *map[int] bool) {
	curSize := len(*fullSet)
	for i:=0; i < size; i++{
		temp := customrandom.GetRandomNumber(0,2147483647)
		(*fullSet)[temp] = true
		for curSize + i + 1 != len(*fullSet){
			temp := customrandom.GetRandomNumber(0,2147483647)
			(*fullSet)[temp] = true
		}
		(*hFunc) = append((*hFunc),temp)
	}
	// println("hashes :")
	// for i:=0; i < size; i++{
	// 	println((*arr)[i])
	// }
	// println("###")
}

func hashFunction( num int) int{
	return (num*(2654435761)>>32)
}

//Recording - acting nth hash function 
func (hFunc *HFunctionCollection)Recording(bitMap *[]bool ,flowID, elementID, m int) {
	idxAmongLBits := hashFunction(elementID)%(len(*hFunc))
	idxAmongMbits := (hashFunction(flowID)^(*hFunc)[idxAmongLBits]) % m
	(*bitMap)[idxAmongMbits] = true
}

//EstimateFlowSpread :
func (hFunc *HFunctionCollection) EstimateFlowSpread(bitMap *[]bool ,flowID, m , totalZeroes int) float64{
	l := len(*hFunc)
	var zeroesInVbitMap float64
	for i:=0; i < l; i++{
		if (*bitMap)[(hashFunction(flowID)^(*hFunc)[i]) % m] == false{
			zeroesInVbitMap++
		}
	}
	return float64(l)*(math.Log(float64(totalZeroes)/float64(m)) - math.Log(zeroesInVbitMap/float64(l)))
}
