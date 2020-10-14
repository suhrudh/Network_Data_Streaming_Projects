package stream

import (
	"project1/customrandom"
)


//GetStream : gets the stream
func GetStream(numOfUniqueFlowIDs int) *map[int]bool{
	set := make(map[int]bool)
	for len(set) < numOfUniqueFlowIDs{
		set[customrandom.GetRandomNumber(0,2147483647)] = true
	}
	return &set
}