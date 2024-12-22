package main

import "fmt"

func swap(ptr1, ptr2 *int) () {
    tmp := *ptr1
    *ptr1 = *ptr2
    *ptr2 = tmp
}

func bubblesort(ts []int) () {
    for size := len(ts) ; size > 1; {
        newsize := 0
        for cpt := 1; cpt < size; cpt ++ {
            if ts[cpt-1] > ts[cpt] {
                swap(&(ts[cpt-1]),&(ts[cpt]))
                newsize = cpt
            } 
        }
        size = newsize
    }
}

func main() {
    unsorted := []int{42,-6,27,10,31,49,1,3}
    fmt.Printf("Avant : ")
    for i := range unsorted {
        fmt.Printf(" %d ", unsorted[i])
    }
    bubblesort(unsorted)
    fmt.Printf("\nAprès : ")
    for i := range unsorted {
        fmt.Printf(" %d ", unsorted[i])
    }
}
