package main
import "fmt"

func fact(n int) int {
    res := 1
    for i := 1; i <= n; i++ {
        res *= i
    }
    return res
}

func perm(n, r int) int {
    return fact(n) / fact(n-r)
}

func comb(n, r int) int {
    return fact(n) / (fact(r) * fact(n-r))
}

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    fmt.Println(perm(a,c), comb(a,c))
    fmt.Println(perm(b,d), comb(b,d))
}
