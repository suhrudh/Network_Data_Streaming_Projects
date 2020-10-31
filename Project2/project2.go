package main

import(
	"os"
	"strconv"
	"project2/hash"
	"project2/set"
	"project2/bloomfilter"
)

func main(){
	//inputs : number of ElementsEncoded, number of Bits In Filter, number of Hashes
	bloomFilterDemo(1000, 10000, 7)

	//inputs : number of Elements Encoded, number of Elements Removed, number of Elements Added, number of Counters, number of Hashes
	countingBloomFilterDemo(1000, 500, 500, 10000, 7)

	//inputs : number of sets, number of elements in each set, number of Bits In Each Filter, number Of Hashes
	codedBloomFilter(7, 1000, 30000, 7)
}

func bloomFilterDemo(numberOfElementsEncoded, numberOfBitsInFilter, numberOfHashes int){
	var hashCollection hash.HFunctionCollection // collection of hash functions
	hashCollection.CreateHashCollection(numberOfHashes)

	var bloom bloomfilter.BloomFilter
	bloom = make([]bool, numberOfBitsInFilter)
	input := set.CreateRandomSet(numberOfElementsEncoded);

	for val := range *input{
		bloom.Encoding(val, &hashCollection)
	}

	count1 := 0
	for val := range *input{
		if bloom.LookUp(val, &hashCollection){
			count1++
		}
	}

	setB := set.GetRandomSetOtherThan(input, numberOfElementsEncoded)

	count2 := 0
	for val := range *setB{
		if bloom.LookUp(val, &hashCollection){
			count2++
		}
	}

	file, err :=  os.Create("output/bloomfilter_output.txt")
	if err != nil{
		return
	}
	defer file.Close()
	file.WriteString("elements found from lookup in A : " + strconv.Itoa(count1)+ "\n")
	file.WriteString("elements found from lookup in B : " + strconv.Itoa(count2)+ "\n")
}

func countingBloomFilterDemo(numOfElementsEncoded, numOfElementsRemoved, numOfElementsAdded, numOfCounters, numOfHashes int){
	var hashCollection hash.HFunctionCollection // collection of hash functions
	hashCollection.CreateHashCollection(numOfHashes)

	var countingBloom bloomfilter.CountingBloomFilter
	countingBloom = make([]int, numOfCounters)
	setA := set.CreateRandomSet(numOfElementsEncoded);
	copySetA := make(map[int] bool)
	for elem := range *setA{
		copySetA[elem] = true
	}

	//encode elements from setA
	for val := range *setA{
		countingBloom.Encoding(val, &hashCollection)
	}

	//form setB splitting setA 
	setB := set.SplitSet(setA,numOfElementsRemoved)

	//remove elements of initial setA from bloom filter
	for val := range *setB{
		countingBloom.RemoveElement(val, &hashCollection)
	}


	//add 500 elements to bloom filter that is not part of initial set
	setC := set.GetRandomSetOtherThan(&copySetA,numOfElementsAdded)
	for val := range *setC{
		countingBloom.Encoding(val, &hashCollection)
	}
	count :=0


	//lookup original elements in the filter
	for val := range copySetA{
		if countingBloom.LookUp(val, &hashCollection){
			count++
		}
	}

	file, err :=  os.Create("output/counting_bloomfilter_output.txt")
	if err != nil{
		return
	}
	defer file.Close()
	file.WriteString("elements found from lookup in A : " + strconv.Itoa(count)+ "\n")
}

func codedBloomFilter(numOfSets, numOfElementsInEachSet, numOfBitsInEachFilter, numOfHashes int){
	setSize := numOfElementsInEachSet
	filterSize := numOfBitsInEachFilter
	g := numOfSets
	numOfhashes := numOfHashes
	// numOfBits := 0
	// for trash > 0{
	// 	trash = trash >> 1
	// 	numOfBits++
	// }

	//intializing Coded Bloom Filter
	var codedBloomFilter bloomfilter.CodedBloomFilter
	codedBloomFilter.Intialization(g, filterSize)

	// create g different sets
	sets := set.CreateAndSplitIntoMultipleSets(setSize,g)

	// 1 to g codes => log2 g hash functions
	var hashCollection hash.HFunctionCollection // collection of hash functions
	hashCollection.CreateHashCollection(numOfhashes)

	for i:= 0; i < g; i++{
		for flowID := range *(*sets)[i]{
			codedBloomFilter.Encoding( i+1, flowID, &hashCollection)
		}
	}
	count := 0
	for i:= 0; i < g; i++{
		for flowID := range *(*sets)[i]{
			
			if codedBloomFilter.LookUp(flowID, &hashCollection) == i+1{
				count++;
			}
		}
	}

	file, err :=  os.Create("output/coded_bloomfilter_output.txt")
	if err != nil{
		return
	}
	defer file.Close()
	file.WriteString("coded bloom filter lookup success : " + strconv.Itoa(count)+ "\n")
}

