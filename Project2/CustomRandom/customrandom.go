package customrandom


import ("math/rand"
		"time")

//GetRandomNumber - gets random number
func GetRandomNumber(start, size int) int{
	seed := rand.NewSource(time.Now().UnixNano())
	randNum := rand.New(seed)

	return start + randNum.Intn(size)  
}



