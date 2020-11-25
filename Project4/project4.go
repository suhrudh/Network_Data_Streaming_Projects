package main

import(
	"project4/preprocessing"
	"project4/hash"
	"project4/set"
	"strconv"
	"os"
)

func main(){
	m := 500000 // bitmap size
	l := 500 //single virtual bitmap size

	bitMap := make([]bool,m)
	inputMap, inputKeys , inputTransformation , _ := preprocessing.Preprocess()

	virtualBitmapLocator := make(map[string] *hash.HFunctionCollection)
	fullSet1 := make(map[int] bool)

	//setting virtual bitmap to bitmap mappings for each flow
	for _,key := range (*inputKeys){
		var tempHFunc hash.HFunctionCollection
		tempHFunc.CreateHashCollection(l, &fullSet1)
		virtualBitmapLocator[key] = &tempHFunc
	}

	flowElementsMap := make(map[string] *[]int)
	fullSet2 := make(map[int] bool)

	//creating distinct elements for each flow
	for _,key := range (*inputKeys){
		flowElementsMap[key] = set.CreateListRandomNumbersUniqueTo(&fullSet2, (*inputMap)[key])
	}

	//endcoding
	for i,key := range (*inputKeys){
		elements := flowElementsMap[key]
		for elementID := range (*elements){
			(*virtualBitmapLocator[key]).Recording(&bitMap, (*inputTransformation)[i], elementID, m)
		}
	}

	//finding total zeroes
	totalZeroes := 0
	for _,entry := range(bitMap){
		if !entry{
			totalZeroes++
		}
	} 

	file1, err :=  os.Create("output/x_values.txt")
	if err != nil{
		return
	}
	file2, err :=  os.Create("output/y_values.txt")
	if err != nil{
		return
	}
	//results
	for i,key := range (*inputKeys){
		estimation := (*virtualBitmapLocator[key]).EstimateFlowSpread(&bitMap, (*inputTransformation)[i], m, totalZeroes)
		//println( strconv.Itoa((*inputMap)[key]) + " : " + strconv.Itoa(int(estimation)))
		file1.WriteString(strconv.Itoa((*inputMap)[key])+"\n")
		file2.WriteString(strconv.Itoa(int(estimation))+"\n")
	}

}








