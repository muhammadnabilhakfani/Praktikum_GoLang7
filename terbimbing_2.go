package main
import "fmt"

func main() {
var x, y, n int
var S float64
fmt.Scan(&x, &y)
n = (y - x) + 1
S = float64(n * (x + y)) / 2
fmt.Println(S)
}
