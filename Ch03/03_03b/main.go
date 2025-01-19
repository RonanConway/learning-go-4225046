package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])

  var numbers =[5]int{5,3,1,2,4}
	fmt.Println(numbers)
	fmt.Println("Number of numbers: ", len(numbers))

	var newcolors = make([]string, 0, 3)
  newcolors = append(newcolors, "Indigo", "Violet", "Puce")
	fmt.Println(newcolors)

	newcolors = append(newcolors, "Black", "Purple")
	fmt.Println(newcolors)

  newcolors = remove(newcolors, 0)
	fmt.Println(newcolors)

	sort.Strings(newcolors)
	fmt.Println(newcolors)

}


func remove(slice [] string, i int) [] string{
	return append(slice[:i], slice[i+1:]...)
}