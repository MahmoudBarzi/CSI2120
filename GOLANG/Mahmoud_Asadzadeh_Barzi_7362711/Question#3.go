package main
import "fmt"
import "math/rand"
func abs(number float32)(results float32){
  if number < 0{
    number = number*-1
  }
  return number
}
func absDiff(sliceA , sliceB []float32)(res []float32, err error){
 	
 if(len(sliceA)==len(sliceB)){
 		for i := range sliceA{
 			res = append (res,abs(sliceA[i]-sliceB[i]))
 		}
    return res,err  
 	}
  if(len(sliceA) != len(sliceB)){
    fmt.Println("Slices are not the same length.")
  }
  return res,err	
}

func RandomArray(len int) []float32 {
	array := make([]float32, len)
	for i := range array {
		array[i] = rand.Float32()
	}
	return array
}
func Process(array []float32, theChannel chan float32){
  var firstHalf []float32
  firstHalf = array[:(len(array)/2)]
  var secondHalf []float32 
  secondHalf = array[(len(array)/2):]
  theSum := 0.0
  theSum32 :=float32(theSum)
  results,_ := absDiff(firstHalf,secondHalf)
  for _,theFloats:=range results{
  	theSum32 += theFloats 
  } 
  theChannel <- theSum32
}
func main(){
	rand.Seed(100) // use this seed value
	out := make(chan float32)
	defer close(out)
	sum64 := 0.0
	sum := float32(sum64)
	for i := 0; i<1000 ; i++ {
		a:= RandomArray(2*(50+rand.Intn(50)))
		go Process(a,out)
	}
	for i := 0; i<1000 ; i++{
		temp:= <-out
		sum = sum+temp
	}
	fmt.Println(sum)
	// ******
	// read here the results of the processing
	// and sum these results
	//fmt.Println(sum)
}































