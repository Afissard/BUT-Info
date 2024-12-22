package main 

import "fmt"

func swap (p1 *int32, p2 *int64) {
    tmp1 := *p1 
    tmp2 := *p2 
    *p2 = int64(tmp1)
    *p1 = int32(tmp2)
}

func main() {
    var a int32 
    var b int64 

    a = 42
    b = int64(0xCAFEDEADBEEF)

    swap(&a,&b)
    fmt.Println(a, b)
}
