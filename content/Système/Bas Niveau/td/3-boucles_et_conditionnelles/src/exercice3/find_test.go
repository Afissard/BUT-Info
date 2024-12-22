package main

import "testing"

func TestFind(t *testing.T) {
    d1 := []int{1, 2, 5, 6, 7, 9, 10}

    found, index := Find(d1, 1)
    if !found || index != 0 {
        t.Logf("Find(d1, %d) = %t %d but want %t %d", 9, found, index, true, 5)
    }

    found, index = Find(d1, 10)
    if !found || index != 6 {
        t.Logf("Find(d1, %d) = %t %d but want %t %d", 10, found, index, true, 6)
    }

    found, index = Find(d1, 5)
    if !found || index != 2 {
        t.Logf("Find(d1, %d) = %t %d but want %t %d", 5, found, index, true, 2)
    }

    found, index = Find(d1, 2)
    if !found || index != 1 {
        t.Logf("Find(d1, %d) = %t %d but want %t %d", 2, found, index, true, 1)
    }

    found, index = Find(d1, 9)
    if !found || index != 5 {
        t.Logf("Find(d1, %d) = %t %d but want %t %d", 9, found, index, true, 5)
    }

    found, index = Find(d1, 0)
    if found {
        t.Logf("Find(d1, %d) = %t %d but want %t", 0, found, index, false)
    }

    found, index = Find(d1, 4)
    if found {
        t.Logf("Find(d1, %d) = %t %d but want %t", 4, found, index, false)
    }

    //found, index = Find(d1, 11)
    //if found {
        //t.Logf("Find(d1, %d) = %t %d but want %t", 11, found, index, false)
    //}
}
