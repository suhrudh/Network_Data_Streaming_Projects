package set

import (
	"project2/customrandom"
)


//CreateRandomSet : creates a set
func CreateRandomSet(setSize int) *map[int]bool{
	set := make(map[int]bool)
	for len(set) < setSize{
		set[customrandom.GetRandomNumber(0,2147483647)] = true
	}
	return &set
}

//GetRandomSetOtherThan : gets random set of desired size not in input set.
func GetRandomSetOtherThan(inputSet *map[int]bool, desiredSize int)*map[int]bool{
	set := make(map[int]bool)
	for len(set) < desiredSize{
		curRandom := customrandom.GetRandomNumber(0,2147483647)
		if _,fl := (*inputSet)[curRandom]; fl{ continue }
		set[curRandom] = true
	}
	return &set
}

//SplitSet : splits a input set by separating a set of desired set from input set.
func SplitSet(inputSet *map[int]bool, desiredSize int)*map[int]bool{
	set := make(map[int]bool)
	temp := make([]int, len(*inputSet))
	for i := range *inputSet{
		temp = append(temp, i)
	}

	for len(set) < desiredSize{
		curRandom := customrandom.GetRandomNumber(0,2147483647)%len(temp)
		set[temp[curRandom]] = true
		delete(*inputSet, temp[curRandom])
		temp[curRandom] = temp[len(temp)-1]
		temp = temp[:len(temp)-1]
	}

	return &set
}

//AddSecondSetToFirst : Add the setA into setB
func AddSecondSetToFirst(setA , setB *map[int]bool){
	for elem := range *setA{
		(*setB)[elem] = true
	}
}

//CreateAndSplitIntoMultipleSets : split into multiple sets
func CreateAndSplitIntoMultipleSets(sizeOfSet ,num int) *[] *map[int] bool{
	//println(num)
	sets := make([]*map[int] bool, num)
	for j:=0 ; j < num; j++{
		temp:= make( map[int] bool)
		sets[j] = &temp
	}
	// println(len(*sets[0]))
	baseSet := CreateRandomSet(num*sizeOfSet)
	i:=0
	for val := range *baseSet{
		//println(i)
		(*(sets[i/sizeOfSet]))[val] = true
		i++
	}
	return &sets
}