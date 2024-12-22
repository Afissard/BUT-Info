package main

import "fmt"

func Find (data []int, elem int) (bool, int) {
    lo, hi := 0, len(data)
    mid := lo + (hi - lo) / 2

    for ; data[mid] != elem && hi > lo; {
        if data[mid] > elem {
            hi = mid - 1
            mid = lo + (hi - lo) / 2
        } else {
            lo = mid + 1
            mid = lo + (hi - lo) / 2
        }
    }
    return data[mid] == elem, mid
}

func main() {
    
    d1 := []int{1, 2, 5, 6, 7, 9, 10}
    found, index := Find(d1, 4)
    if found {
        fmt.Printf("Find(d1, %d) = %t %d but want %t", 0, found, index, false)
    }
}
