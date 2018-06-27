package fizzbuzz

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
    f, _ := strconv.ParseFloat(os.Args[1], 64)
    var describer Describer
    describer = ChosenNumber(f) 
    result := describer.Describe()
    fmt.Println(result)
}
