package main
import (
	"fmt"
	"math"
)
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x1-x2, y1-y2)
}
func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 := diDalam(cx1, cy1, r1, x, y)
	in2 := diDalam(cx2, cy2, r2, x, y)

	switch {
	case in1 && in2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case in1:
		fmt.Println("Titik di dalam lingkaran 1")
	case in2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
