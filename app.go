package main

import "fmt"
import "flag"
import "os"
func main() {
    calc := flag.NewFlagSet("calc", flag.ExitOnError)
    doubledNum := calc.Int("double",0 , "Doubles given number.")
    if len(os.Args) < 2 {
        fmt.Println("usage: calc [args]")
        os.Exit(1)
    }
    switch os.Args[1] {
        case "calc":
            calc.Parse(os.Args[2:])
            doubleMe(*doubledNum)
            calc.Args()
    }
}
func doubleMe(num int) {
    if num==0 {
        print()
    } else {
        fmt.Println(num*2)
    }
}