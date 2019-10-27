package main

import "fmt"
import "flag"
func main() {
    doubledNum := flag.Int("double",0 , "Doubles given number.")
	flag.Parse()
	doubleMe(*doubledNum)
	flag.Args()
}
func doubleMe(num int) {
    if num==0 {
        print()
    } else {
        fmt.Println(num,"* 2 is",num*2)
    }
}