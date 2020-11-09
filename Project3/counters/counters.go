package counters

import (
	"project3/hash"
	"project3/utility"
	"sort"
)

// //PseudoMsb : no, of bits to shift to look at MSB
// var PseudoMsb = 0

//CountMin array of counts
type CountMin [][]int

//InitializeCountMin : Initializes CountMin
func InitializeCountMin(k, w int) *CountMin {
	var temp CountMin
	temp = make([][]int, k, k)
	for i := range temp {
		temp[i] = make([]int, w, w)
	}
	return &temp
}

//Recording : records the flowIDs
func (arr *CountMin) Recording(flowID, count int, h *hash.HFunctionCollection) {
	w := len((*arr)[0])

	for i := range *arr {
		(*arr)[i][(*h).ActHashN(i, flowID, w)] += count
	}
}

//Querying :  query's the minimum count
func (arr *CountMin) Querying(flowID int, h *hash.HFunctionCollection) int {
	count := 2147483646
	w := len((*arr)[0])
	for i := range *arr {
		count = min(count, (*arr)[i][(*h).ActHashN(i, flowID, w)])
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//CounterSketch : storage
type CounterSketch [][]int

//Recording : records the flowIDs in counter sketch manner
func (arr *CounterSketch) Recording(flowID, count int, h *hash.HFunctionCollection) {
	w := len((*arr)[0])
	for i := range *arr {
		hashGenerated := (*h).ActHashNWithoutMod(i, flowID)
		if (hashGenerated >> 30)%2 == 0 {
			(*arr)[i][(hashGenerated)%w] -= count
		}else{
			(*arr)[i][hashGenerated%w] += count
		}
	}
	
}

// //PseudoMSBextractor : MSB extractor
// func PseudoMSBextractor(w int) int {
// 	temp := 1
// 	for w > 0 {
// 		w = w >> 1
// 		temp++
// 	}

// 	return temp
// }

//Querying :  query's median of counts
func (arr *CounterSketch) Querying(flowID int, h *hash.HFunctionCollection) int {
	w := len((*arr)[0])
	var temp []int
	for i := range *arr {
		temp = append(temp, (*arr)[i][((*h).ActHashNWithoutMod(i, flowID))%w] )
	}
	sort.Ints(temp)
	if(len(temp)%2 == 0){
		temp[len(temp)/2] += temp[(len(temp)/2)-1]
	}
	return utility.Abs(temp[len(temp)/2])
}

//InitializeCounterSketch : Initializes CountMin
func InitializeCounterSketch(k, w int) *CounterSketch {
	var temp CounterSketch
	temp = make([][]int, k, k)
	for i := range temp {
		temp[i] = make([]int, w, w)
	}
	return &temp
}
