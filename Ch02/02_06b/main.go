package main

import (
	"fmt"
	"time"
)

func main() {
  t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("GO launched at %s \n", t)

	n := time.Now()
	fmt.Printf("The time currently is %s \n", n)
	fmt.Printf("The objects type is %T\n", n)
  fmt.Printf(n.Format(time.ANSIC) + "\n")

	tomrrow := n.AddDate(0, 0, 1)
  fmt.Printf(tomrrow.Format(time.ANSIC) + "\n")
  format := "Mon 2006-02-01"
  fmt.Printf(tomrrow.Format(format) + "\n")

}
