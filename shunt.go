package main

import (
    "fmt"
)

func intopost(infix string) string {
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}

    

    pofix, s := []rune{}, []rune{} //s = stack
    return string(pofix)
}

func main() {

    // Answer: ab.c
    fmt.Println("infix: ", "a.b.c*")
    fmt.Println("postfix: ", intopost("a.b.c*"))

    // Answer: abd|.*
    fmt.Println("infix: ", "(a.(b|d))*")
    fmt.Println("postfix: ", intopost("(a.(b|d))"))

    // Answer: abd|.c*.
    fmt.Println("infix ", "a.(b|d).c*")
    fmt.Println("postfix ", intopost("a.(b|d).c*"))
 
    //Answer: abb.+.c.
    fmt.Println("infix ", "a.(b.b)+.c")
    fmt.Println("postfix ", intopost("a.(b.b)+.c"))
}