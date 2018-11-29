package check

import (
	"fmt"
	"math"
	"sort"
)

var i int = 0
var ds float64
var s float64
var hd float64 = math.Pi / 180
var du float64 = 180 / math.Pi


func comment(ds float64, e float64, s float64, a float64) float64 {
	return math.Sqrt((math.Pow(e, 2)) + math.Pow(((ds-e)/(math.Tan(a)))-s, 2))
}

func getInfo(begin float64, end float64) (length int, f float64, r []float64) {
	length = int(end-begin) + 1
	f = end - begin
	r = make([]float64, length)
	return
}

func sendData(r []float64, length int) float64 {
	sort.Float64s(r)
	fmt.Println("Max:", r[length-1])
	i = 0
	return r[length-1]
}
