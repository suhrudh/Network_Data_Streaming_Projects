package bloomfilter

import("project2/hash")

//BloomFilter : Bloom
type BloomFilter []bool

//Encoding field ids into bloom filter
func (bloom *BloomFilter) Encoding(flowID int, h *hash.HFunctionCollection ){
	for i:= 0; i < len(*h); i++{
		(*bloom)[(*h).ActHashN(i, flowID, len(*bloom))] = true
	}
}

//LookUp a flowID in BloomFilter
func (bloom *BloomFilter) LookUp(flowID int, h *hash.HFunctionCollection) bool{
	flag := true
	for i:= 0; i < len(*h); i++{
		if !((*bloom)[(*h).ActHashN(i, flowID, len(*bloom))]){
			flag = false
			break
		}
	}
	return flag
}



//CountingBloomFilter : Bloom
type CountingBloomFilter []int

//Encoding field ids into bloom filter
func (bloom *CountingBloomFilter) Encoding(flowID int, h *hash.HFunctionCollection ){
	for i:= 0; i < len(*h); i++{
		(*bloom)[(*h).ActHashN(i, flowID, len(*bloom))]++
	}
}

//LookUp a flowID in BloomFilter
func (bloom *CountingBloomFilter) LookUp(flowID int, h *hash.HFunctionCollection) bool{
	flag := true
	for i:= 0; i < len(*h); i++{
		if (*bloom)[(*h).ActHashN(i, flowID, len(*bloom))] <= 0{
			flag = false
			break
		}
	}
	return flag
}

//RemoveElement a flowID in BloomFilter
func (bloom *CountingBloomFilter) RemoveElement(flowID int, h *hash.HFunctionCollection) {
	for i:= 0; i < len(*h); i++{
		(*bloom)[(*h).ActHashN(i, flowID, len(*bloom))]--
	}
}

//CodedBloomFilter : implementation
type CodedBloomFilter [] BloomFilter

//Intialization : intializinga coded bloom filter
func (codedBloom *CodedBloomFilter) Intialization(numOfGroups , filterSize int){
	var temp []BloomFilter
	for numOfGroups > 0{
		temp = append(temp, make([]bool, filterSize))
		numOfGroups = numOfGroups >> 1
	}
	*codedBloom = temp;
}

//Encoding : encoding elements into coded bloom filter
func (codedBloom *CodedBloomFilter) Encoding(setNumber ,flowID int, h *hash.HFunctionCollection){
	i := 0
	for setNumber > 0{
		if setNumber%2 == 1{
			(*codedBloom)[i].Encoding(flowID, h)
		}
		setNumber = setNumber >> 1
		i++;
	}
}

//LookUp a flowID in BloomFilter
func (codedBloom *CodedBloomFilter) LookUp(flowID int, h *hash.HFunctionCollection) int{
	code := 0
	
	for i := len(*codedBloom)-1; i >= 0 ; i--{
		if (*codedBloom)[i].LookUp(flowID, h){
			code++;
		}
		if i == 0 {break}
		code = code << 1
	}
	//println(code)
	return code
}