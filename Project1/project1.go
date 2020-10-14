package main

import(
	//"fmt"
	"strconv"
	"project1/stream"
	"project1/hash"
	"os"
)

func main(){
	//inputs : number of table entries, number of flows, number of hashes
	MultiHashDemo(1000, 1000, 3)
	//inputs : number of table entries, number of flows, number of hashes, number of cuckoo steps
	Cuckoo(1000,1000, 3, 2)
	//inputs : number of table entries, number of flows, d value
	DLeftHashingDemo(1000,1000,4)
}

//DLeftHashingDemo , numOfTableEntries, numOfFlows, d int
func DLeftHashingDemo(numOfTableEntries, numOfFlows, d int){
	stream := stream.GetStream(numOfFlows) // input stream
	storage := make([]int, numOfTableEntries) //hash table
	for i:= 0; i < numOfTableEntries ; i++{
		storage[i] = -1
	}

	var hashCollection hash.HFunctionCollection // collection of hash functions, Declaration
	hashCollection.CreateHashCollection(d) // Initializing

	//perfomring Dleft Hashing for each element in the stream
	for flowID := range *stream{
		hashCollection.DleftHashing(&storage, flowID)
	}

	//counting filled entries
	var count int
	for i:=0; i < len(storage); i++{
		if storage[i] != -1{
			count++
		}
		//fmt.Println(storage[i])
	}

	printOutput(count, "dleft", &storage)
}

func printOutput(count int, hashingType string, storage *[]int){

	file, err :=  os.Create("output/"+ hashingType+ "_output.txt")
	if err != nil{
		return
	}
	defer file.Close()
	file.WriteString("number of entries filled in "+ hashingType +" :" + strconv.Itoa(count)+ "\n")

	//printing table entries
	for i:=0; i < len(*storage); i++{
		if (*storage)[i] < 0 { (*storage)[i] = 0}
		file.WriteString( strconv.Itoa((*storage)[i])+ "\n")
	}
}

//Cuckoo : numOfTableEntries, numOfFlows, numOfHashes, steps int
func Cuckoo(numOfTableEntries, numOfFlows, numOfHashes, steps int){
	stream := stream.GetStream(numOfFlows) // input stream
	storage := make([]int, numOfTableEntries) //hash table
	for i:= 0; i < numOfTableEntries ; i++{
		storage[i] = -1
	}

	var hashCollection hash.HFunctionCollection // collection of hash functions, declaration
	hashCollection.CreateHashCollection(3) // intialization

	//performing cuckoo hashing for each element in the stream
	for flowID := range *stream{
		hashCollection.CuckooHashing(&storage, flowID,steps)
	}

	//counting filled entries
	var count int
	for i:=0; i < len(storage); i++{
		if storage[i] != -1{
			count++
		}
		//fmt.Println(storage[i])
	}

	printOutput(count, "cuckoo", &storage)
}

//MultiHashDemo : numOfTableEntries, numOfFlows, numOfHashes int
func MultiHashDemo(numOfTableEntries, numOfFlows, numOfHashes int){
	stream := stream.GetStream(numOfFlows) // input stream
	storage := make([]int, numOfTableEntries) //hash table
	for i:= 0; i < numOfTableEntries ; i++{
		storage[i] = -1
	} 
	// for i:= 0; i < 1000 ; i++{
	// 	fmt.Println(storage[i]) //storage[i] = -1
	// } 
	var hashCollection hash.HFunctionCollection // collection of hash functions
	hashCollection.CreateHashCollection(numOfHashes)

	//performing multi hash for each element in the stream
	for flowID := range *stream{
		hashCollection.MultiHashing(&storage, flowID)
	}

	//counting filled entries
	var count int
	for i:=0; i < len(storage); i++{
		if storage[i] != -1{
			count++
		}
		//fmt.Println(storage[i])
	}

	printOutput(count, "multi", &storage)
}