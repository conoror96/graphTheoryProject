//Shunting yard algorithm.
//This is a method for parsing mathematical expressions specified in infix notation. It produces a postfix notation string.
//Ref: https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
package main

import (
    "fmt"
)

func intopost(infix string) string {
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}

    pofix, s := []rune{}, []rune{} //s = stack

     //loops through infix, returns index of char we're currently reading. r = character. 
     // when call range on infix converts each element of the string to a rune
     // range connstruct on a string converts the string to an array of runes using UTF8
    for _, r:= range infix {
        switch {
        case r == '(': //if open bracket, stick it onto end of the stack temporarily.
            s = append(s, r) 
        case r == ')': //if closing bracket, pop thing  off stack until open bracket is found. 
            for s[len(s)-1] != '(' {
                pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] // takes the last element of the stack, puts it onto last element of pofix
            }
            s = s[:len(s)-1]
        case specials[r] > 0: // if not a special character = 0 or NULL
            for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
                pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
            }
            s = append(s, r)
        default:
          pofix = append(pofix, r)
        }
    }

    //if theres anything left on top of stack, append onto output
    for len(s) > 0 {
        pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
    }


    return string(pofix)
}

func main() {

    // Answer: ab.c
    fmt.Println("infix: ", "a.b.c*")
    fmt.Println("postfix: ", intopost("a.b.c*"))

    // Answer: abd|.*
    fmt.Println("infix: ", "(a.(b|d))*")
    fmt.Println("postfix: ", intopost("(a.(b|d))*"))

    // Answer: abd|.c*.
    fmt.Println("infix ", "a.(b|d).c*")
    fmt.Println("postfix ", intopost("a.(b|d).c*"))
 
    //Answer: abb.+.c.
    fmt.Println("infix ", "a.(b.b)+.c")
    fmt.Println("postfix ", intopost("a.(b.b)+.c"))
}