package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
  if len(os.Args) > 1 {
    f, _ := strconv.ParseFloat(os.Args[1], 62)
    numberToOutput(f)
  } else {
	  for number := 1.0; number <= 100; number++ {
		  numberToOutput(number)
	  }
  }
}

func numberToOutput(number float64) {
  var describer Describer
	describer = ChosenNumber(number)
	result := describer.Describe()
	fmt.Println(result)
}
