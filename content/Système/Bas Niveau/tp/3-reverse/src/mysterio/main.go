package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func GetInput(in *bufio.Reader) (string) {

    s, err := in.ReadString('\n')
    if err != nil {
        log.Fatal("Cannot read input:", err)
    }
    return strings.TrimSpace(s)
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("Enter your answer for step 1 > ")
    answer := GetInput(reader)
    if !Step1(answer) {
        fmt.Println("You failed at step 1!")
        os.Exit(0)
    }

    fmt.Println("Nicely done. Will you also pass step 2?")
    fmt.Printf("Enter your answer for step 2 > ")
    answer = GetInput(reader)
    if !Step2(answer) {
        fmt.Println("Apparently your luck as run out! Failure at step 2.")
        os.Exit(0)
    }

    fmt.Println("Good job. Let's try step 3?")
    fmt.Printf("Enter your answer for step 3 > ")
    answer = GetInput(reader)
    if !Step3(answer) {
        fmt.Println("You failed at step 3!")
        os.Exit(0)
    }

    fmt.Println("You've been lucky so far, but trust me, you'll never pass step 4.")
    fmt.Printf("Enter your answer for step 4 > ")
    answer = GetInput(reader)
    if !Step4(answer) {
        fmt.Println("I told you. Step 4 is too difficult for you")
        os.Exit(0)
    }

    fmt.Println("GG, WP")
    os.Exit(0)
}

