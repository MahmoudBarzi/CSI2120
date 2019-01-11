package main
import "fmt"
type Bread struct{
	name string 
	ingredients map[string]Item
	weight float32
	Baking
}
type Baking struct{
		bakeTime int
		coolTime int
		temperature int
}
type Item struct{
	weight int
}
type Baker interface{
	shoppingList(listOfIngredients map[string]Item) (listOfIngredients1,listOfIngredients2 map[string]Item)
	printBakeInstructions()
	printBreadInfo()
}
func newBread()(*Bread){
	wholeWheat := new(Bread)
	wholeWheat.name = "Whole Wheat bread"
	wholeWheat.ingredients = make(map[string]Item)
	wholeWheat.ingredients["whole wheat flour"] = Item{500}
	wholeWheat.ingredients["yeast"] = Item{25}
	wholeWheat.ingredients["salt"] = Item{25}
	wholeWheat.ingredients["sugar"] = Item{50}
	wholeWheat.ingredients["butter"] = Item{50}
	wholeWheat.ingredients["water"] = Item{350}
	wholeWheat.weight = 1
	wholeWheat.Baking.bakeTime = 120
	wholeWheat.Baking.coolTime = 60
	wholeWheat.Baking.temperature = 180
	return wholeWheat
}
func NewBreadVariation (name string, ingredientsToBeAdded, ingredientsToBeRemoved map[string]Item) (*Bread){
	theBread := newBread()
	theBread.name = name
	for k1,_ :=range theBread.ingredients{
		for k2,_:=range ingredientsToBeRemoved{
			if k1 == k2{
				delete(theBread.ingredients,k2)
			}
		}
	}
	for k,v :=range ingredientsToBeAdded{
			theBread.ingredients[k] = v
	}
	return theBread
}
func(bread *Bread) shoppingList(listOfIngredients map[string]Item) (needed,leftOver map[string]Item){
	needed = make(map[string]Item)
	leftOver = make(map[string]Item)
	for k1,v1 := range bread.ingredients {
			flag := false
			for k2,_ := range listOfIngredients {
				 if k1 == k2 {
				// 	nameVal1 := bread.ingredients[k1]
				// 	nameVal2 := listOfIngredients[k2]
				// 	nv1 := nameVal1.weight
				// 	nv2 := nameVal2.weight
					leftOver[k1]= Item{bread.ingredients[k1].weight - listOfIngredients[k2].weight}
					flag = true
					break
				}
			}
			if !flag {
				needed[k1] = v1
			}
	}
	return needed,leftOver
}
func(bread *Bread)printBakeInstructions(){
	fmt.Println("Bake at",bread.Baking.temperature,"Celsius for",bread.Baking.bakeTime,"minutes and let cool for",bread.Baking.coolTime,"minutes.")
}
func(bread *Bread)printBreadInfo(){
	fmt.Println(bread.name)
	fmt.Println(bread.ingredients)
	fmt.Println("Weight",bread.weight,"kg")
 }
func main(){
	wholeWheatBread := newBread()
	kitchenMap :=make(map[string]Item)
	kitchenMap["whole wheat flour"] = Item{5000}
	kitchenMap["salt"] = Item{500}
	kitchenMap["sugar"] = Item{1000}
	neededWWB,leftOver := wholeWheatBread.shoppingList(kitchenMap)
	toBeAdded :=map[string]Item{"white flour":{200},"sesame":{50},"whole wheat flour":{250}}
	toBeRemoved :=map[string]Item{"whole wheat flour":{500}}
	sesameBread := NewBreadVariation("Sesame bread",toBeAdded,toBeRemoved)
	neededSB,leftOver := sesameBread.shoppingList(leftOver)
	shoppingList:= make(map[string]Item)
	for k1,v1 := range neededSB{
		flag:= false
		for k2,v2 := range neededWWB{
			if k1 == k2 {
				shoppingList[k1] = Item{v1.weight+v2.weight}
				flag = true
				break
			}
		}
		if !flag{
			shoppingList[k1]= v1
		}
	}
	wholeWheatBread.printBreadInfo()
	fmt.Println("\n")
	sesameBread.printBreadInfo()
	fmt.Println("\n")
	fmt.Println("Shopping List :")
	fmt.Println(shoppingList)
	fmt.Println("\n")
	fmt.Println("Baking Instructions:")
	wholeWheatBread.printBakeInstructions()
	sesameBread.printBakeInstructions()
}














