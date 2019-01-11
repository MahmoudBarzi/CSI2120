package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
func AbsDiff(sliceA , sliceB []float32,version int)(res []float32, err error){

  if(version == 0){
    res,err = AbsDiffVersionA(sliceA , sliceB)
  }
  if(version == -1){
     res,err = AbsDiffVersionB(sliceA , sliceB)
  }
  if(version == 1){
     res,err = AbsDiffVersionC(sliceA , sliceB)
  }
  return res,err
}
//-------------------------------------------------------------------------------
func AbsDiffVersionA(sliceA , sliceB []float32)(res []float32, err error){
 	
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
//-------------------------------------------------------------------------------
func AbsDiffVersionB(sliceA , sliceB []float32)(res []float32, err error){
  
  if(len(sliceA)==len(sliceB)){// both slices are equal length
    for i := range sliceA{
      res = append (res,abs(sliceA[i]-sliceB[i]))
    }
    return res,err  
  }
  if(len(sliceA) > len(sliceB)){//A is the bigger slice
    for i := range sliceA{
      if ( i >= len(sliceB)) {
      res = append (res,abs(sliceA[i] - 0))
      }else{
        res = append (res,abs(sliceA[i]-sliceB[i]))
      } 
    }
  return res,err 
  }
  if(len(sliceA) < len(sliceB)){//B is the bigger slice 
    for i := range sliceB{
      if(i >= len(sliceA)){
        res = append (res,abs(0 - sliceB[i]))
      }else{
        res = append (res,abs(sliceA[i]-sliceB[i]))
      }
    }
  return res,err 
  }
  return res,err  
 }
 //-------------------------------------------------------------------------------
 func AbsDiffVersionC(sliceA , sliceB []float32)(res []float32, err error){
  
  if(len(sliceA)==len(sliceB)){//if both slices have the same length
    for i := range sliceA{
      res = append (res,abs(sliceA[i]-sliceB[i]))
    }
    return res,err  
  }
  if(len(sliceA) < len(sliceB)){//if A has a shorter length
    for i := range sliceA{
      res = append (res,abs(sliceA[i]-sliceB[i]))
    }
    return res,err  
  }
  if(len(sliceA) > len(sliceB)){//if B has a shorter length
     for i := range sliceB{
      res = append (res,abs(sliceA[i]-sliceB[i]))
    }
    return res,err  
  }
  return res,err 
}
//-------------------------------------------------------------------------------
func floatMaker(text string)(results float64){
  results,err := strconv.ParseFloat(text,64)

  if err != nil{
    fmt.Println("This Ain't A Number!")
  }
  return results
}

func abs(number float32)(results float32){
  if number < 0{
    number = number*-1
  }
  return number
}
func main(){
    reader1 := bufio.NewReader(os.Stdin)
    fmt.Print("Previous slice: ")
    text1, _ := reader1.ReadString('\n')
    removeOpenBracket1 := strings.Replace(text1, "[", "", -1)
    removeCloseBracket1:= strings.Replace(removeOpenBracket1, "]", "", -1)
    text1 = removeCloseBracket1    
    floats1 := strings.Fields(text1)
    slice1:=[]float64{} 
    slice1float32 :=[]float32{}
    for i, _:= range floats1{
      slice1 = append(slice1,(floatMaker(floats1[i])))
      slice1float32 = append(slice1float32,float32(slice1[i]))
    }
    reader2 := bufio.NewReader(os.Stdin)
    fmt.Print("Enter another slice of floating point numbers (Anything else to end slice) ")
    text2, _ := reader2.ReadString('\n')
    removeOpenBracket2 := strings.Replace(text2, "[", "", -1)
    removeCloseBracket2:= strings.Replace(removeOpenBracket2, "]", "", -1)
    text2 = removeCloseBracket2 
    floats2 := strings.Fields(text2)
    slice2:=[]float64{} 
    slice2float32:=[]float32{}
    for i, _:= range floats2{
      slice2 = append(slice2,(floatMaker(floats2[i])))
      slice2float32 = append(slice2float32,float32(slice2[i]))
    }
  for{
    Results,_:=AbsDiff(slice1float32,slice2float32,0)//************************
    fmt.Println("Result: ",Results)
    reader3 := bufio.NewReader(os.Stdin)
    fmt.Print("q to quit (Anything else to continue):")
    text3, _ := reader3.ReadString('\n')
    text3 = strings.TrimRight(text3,"\n")
    if strings.Compare(text3, "q") == 0{
      os.Exit(1)
    }
    slice1float32 = slice2float32
    fmt.Println("Previous slice: ",slice1float32)
    reader4:= bufio.NewReader(os.Stdin)
    fmt.Print("Enter another slice of floating point numbers (Anything else to end slice) ")
    text4, _ := reader4.ReadString('\n')
    removeOpenBracket4 := strings.Replace(text4, "[", "", +1)
    removeCloseBracket4:= strings.Replace(removeOpenBracket4, "]", "", -1)
    text4 = removeCloseBracket4 
    floats4 := strings.Fields(text4)
    slice4:=[]float64{} 
    slice2float32 = []float32{}
    for i, _:= range floats4{
      slice4 = append(slice4,(floatMaker(floats4[i])))
      slice2float32 = append(slice2float32,float32(slice4[i]))
    }
  }	
}







































