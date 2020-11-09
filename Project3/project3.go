package main

import(
	"project3/preprocessing"
	"project3/counters"
	"project3/hash"
	"project3/customrandom"
	"project3/utility"
	"os"
	"strconv"
	"sort"
)

func main(){
	//inputs : k , w
	countMinDemo(3, 3000)
	//inputs : k , w
	counterSketchDemo(3, 3000)
	//input : number to be acheived
	activeCounterDemo(1000000)
}

func countMinDemo(k, w int){
	inputMap, inputKeys , inputTransformation , numberOfFlows := preprocessing.Preprocess()
	
	var hashCollection hash.HFunctionCollection // collection of hash functions
	hashCollection.CreateHashCollection(k)

	countMin := counters.InitializeCountMin(k, w)
	file, err :=  os.Create("output/count_min_output.txt")
	if err != nil{
		return
	}

	// recording values
	for i := range *inputKeys{
		(*countMin).Recording( (*inputTransformation)[i], (*inputMap)[(*inputKeys)[i]],&hashCollection)
	}
	outputFlowsizeEstimation := make(map[string] int)
	error := 0
	// flow size estimation
	for i := range *inputKeys{
		outputFlowsizeEstimation[(*inputKeys)[i]] = (*countMin).Querying( (*inputTransformation)[i],&hashCollection)
		error += outputFlowsizeEstimation[(*inputKeys)[i]] - (*inputMap)[(*inputKeys)[i]]
	}

	//sorting by estimated flow size
	sort.SliceStable( (*inputKeys)[:], func(s1 , s2 int) bool{
		return outputFlowsizeEstimation[(*inputKeys)[s1]] > outputFlowsizeEstimation[(*inputKeys)[s2]]
	})

	file.WriteString("Average Error : "+ strconv.Itoa(error/numberOfFlows) + "\n")
	file.WriteString("( Flowid , Estimated flow Size , actual flow size ) \n")
	for i := range (*inputKeys){
		if i == 100{ break}
		file.WriteString("( "+ (*inputKeys)[i] + " , "+strconv.Itoa(outputFlowsizeEstimation[(*inputKeys)[i]])+ " , "+ strconv.Itoa((*inputMap)[(*inputKeys)[i]]) + " )\n")
	}
}

func counterSketchDemo(k, w int){
	inputMap, inputKeys , inputTransformation , numberOfFlows := preprocessing.Preprocess()
	
	var hashCollection hash.HFunctionCollection // collection of hash functions
	hashCollection.CreateHashCollection(k)

	counterSketch := counters.InitializeCounterSketch(k, w)
	file, err :=  os.Create("output/counter_sketch_output.txt")
	if err != nil{
		return
	}
	// counters.PseudoMsb = counters.PseudoMSBextractor(w)

	// recording values
	for i := range *inputKeys{
		(*counterSketch).Recording( (*inputTransformation)[i], (*inputMap)[(*inputKeys)[i]],&hashCollection)
	}
	outputFlowsizeEstimation := make(map[string] int)
	error := 0
	// flow size estimation
	
	
	for i := range *inputKeys{
		outputFlowsizeEstimation[(*inputKeys)[i]] = (*counterSketch).Querying( (*inputTransformation)[i],&hashCollection)
		error +=  utility.Abs(outputFlowsizeEstimation[(*inputKeys)[i]] - (*inputMap)[(*inputKeys)[i]])
	}

	//sorting by estimated flow size
	sort.SliceStable( (*inputKeys)[:], func(s1 , s2 int) bool{
		return outputFlowsizeEstimation[(*inputKeys)[s1]] > outputFlowsizeEstimation[(*inputKeys)[s2]]
	})

	file.WriteString("Average Error : "+ strconv.Itoa(error/numberOfFlows) + "\n")
	file.WriteString("( Flowid , Estimated flow Size , actual flow size ) \n")
	for i := range (*inputKeys){
		if i == 100{ break}
		file.WriteString("( "+ (*inputKeys)[i] + " , "+strconv.Itoa(outputFlowsizeEstimation[(*inputKeys)[i]])+ " , "+ strconv.Itoa((*inputMap)[(*inputKeys)[i]]) + " )\n")
	}
}


func activeCounterDemo(count int){
	var number, exponent uint16
	file, err :=  os.Create("output/active_counter_output.txt")
	if err != nil{
		return
	}
	for count > 0{
		activeIncrease(&number, &exponent)
		count--
	}

	ans := int(number) * pow(exponent)
	file.WriteString(strconv.Itoa(ans))
}

func activeIncrease(number , exponent *uint16){
	if decisionWithProbability((*exponent)){
		(*number)++
		if (*number) == 0{
			(*exponent)++
			(*number) = 32768
		}
	}
}

//this function gives true with a probability of 2^(exponent)
func decisionWithProbability(exponent uint16) bool{
	if exponent == 0 { return true }
	temp := pow(exponent)
	return customrandom.GetRandomNumber(0,2147483648)%temp == 0
}

func pow( num uint16) int{
	temp := 1
	temp = temp << num
	return temp
}