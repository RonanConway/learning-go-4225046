package main

import (
	"fmt"
	// "math"
	"strconv"
)

func main() {
// 	i1, i2, i3 := 12, 45, 68
// 	intSum := i1 + i2 + i3
// 	fmt.Println("Integer Sum: ", intSum)

// f1, f2, f3 := 23.5, 65.1, 76.3
// floatSum := f1 + f2 + f3
// fmt.Println("Float Sum: ", floatSum)

// total := float64(i1) + f2
// fmt.Println("Total Sum: ", total)

// product := float64(i1) * f2
// fmt.Println("Product: ", product)



// sum := math.Round(floatSum*100) / 100

// fmt.Printf("The sum is now %v \n\n", sum)

// fmt.Println("The value of PI is %v", math.Pi)

// circleRadius := 15.5
// circumference := circleRadius * 2 * math.Pi
// fmt.Printf("Circumeference: %.2f\n", circumference)

calculate("10", "5.5")

}



func calculate(value1 string, value2 string){

    flt1, _ := strconv.ParseFloat(value1, 1)
		flt2, _ := strconv.ParseFloat(value2, 1)
    sum := flt1 + flt2
		fmt.Printf("The sum of %v and %v is : %v", flt1, flt2, sum)
}
