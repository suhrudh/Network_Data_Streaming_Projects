package preprocessing

import(
	"strings"
	"io/ioutil"
	"strconv"
	"project4/customrandom"
)

//Preprocess : takes input from the text file
func Preprocess() (*map[string] int, *[]string, *[] int, int){
	f, err := ioutil.ReadFile("project4input.txt")


	inputMap := make(map[string] int)
	var inputTransFormation []int 
	var inputKeys []string
	if err != nil{
		println("failed")
		return &inputMap, &inputKeys, &inputTransFormation, 0 
	}

	temp := strings.Split(string(f), "\n")
	numberOfFlows , _ := strconv.Atoi(strings.Fields(temp[0])[0]);
	//println(numberOfFlows)
	for i, tmp := range temp{
		if i == numberOfFlows{ break}
		if i == 0 { continue }
		temp2 := strings.Fields(tmp)
		//println(tmp)
		inputMap[temp2[0]] = 0
		if len(temp2) == 2{
			inputMap[temp2[0]], _ = strconv.Atoi(temp2[1])
		}
		
		inputKeys =  append(inputKeys, temp2[0])
	}

	for range inputKeys{
		inputTransFormation = append(inputTransFormation, customrandom.GetRandomNumber(0,2147483647))
	}	

	return &inputMap , &inputKeys, &inputTransFormation ,numberOfFlows
}
