package hash

import(
		"project1/customrandom")

//HFunctionCollection , array of random numbers which is XORed with hash function output
type HFunctionCollection []int

//CreateHashCollection , creates collection of hash functions
func (arr *HFunctionCollection)CreateHashCollection(size  int) {

	for i:=0; i < size; i++{
		num := customrandom.GetRandomNumber(0,2147483647)
		if(i > 0 && num == (*arr)[i-1]){
			i--;
		}else{
			(*arr) = append((*arr),customrandom.GetRandomNumber(0,2147483647))
		}
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

//MultiHashing - performing multi hashing on a flowID
func (arr *HFunctionCollection) MultiHashing(storage *[]int , flowID int){
	for i:= 0; i < len(*arr); i++{
		pos := (*arr).ActHashN(i, flowID, len(*storage))
		if (*storage)[pos] == -1{
			(*storage)[pos] = flowID
			break;
		}
	}
}

//DleftHashing - performing D-left hashing on a flowID
func (arr *HFunctionCollection) DleftHashing(storage *[]int , flowID int, ){
	chunk := len(*storage)/len(*arr)
	for i:= 0; i < len(*arr); i++{
		pos := i*chunk + (*arr).ActHashN(i, flowID, chunk)
		if (*storage)[pos] == -1{
			(*storage)[pos] = flowID
			return;
		}
	}
}

//CuckooHashing - cuckoo hashing 
func (arr *HFunctionCollection) CuckooHashing(storage *[]int , flowID, steps int){
	for i:= 0; i < len(*arr); i++{
		pos := (*arr).ActHashN(i, flowID, len(*storage))
		if (*storage)[pos] == -1{
			(*storage)[pos] = flowID
			return;
		}
	}

	for i:= 0; i < len(*arr); i++{
		pos := (*arr).ActHashN(i, flowID, len(*storage))
		if( arr.move(storage, pos, steps-1) ){
			(*storage)[pos] = flowID
			break;
		}
	}
}

//move function moves the current element to a possible empty entry.
func (arr *HFunctionCollection)move(storage *[]int ,conflictIndex , stepsLeft int) bool{
	if stepsLeft == -1{ return false}
	toBeMovedFlowID := (*storage)[conflictIndex]

	for i:=0; i < len(*arr); i++{
		pos := (*arr).ActHashN(i, toBeMovedFlowID, len(*storage))
		if(pos != conflictIndex && (*storage)[pos] == -1){
			(*storage)[pos] = toBeMovedFlowID
			(*storage)[conflictIndex] = -1
			return true
		}
	}

	for i:=0; i < len(*arr); i++{
		pos := (*arr).ActHashN(i, toBeMovedFlowID, len(*storage))
		if(pos != conflictIndex && arr.move(storage, pos, stepsLeft-1)){
			(*storage)[pos] = toBeMovedFlowID
			(*storage)[conflictIndex] = -1
			return true
		}
	}

	return false
}

