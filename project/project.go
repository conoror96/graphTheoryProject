package main

import(
    "fmt"
)

//intopost function from shunt.go
//converts infix notation to postfix notation.
func intopost(infix string) string {
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}

    pofix, s := []rune{}, []rune{} 

    for _, r:= range infix {
        switch {
        case r == '(': 
            s = append(s, r) 
        case r == ')': 
            for s[len(s)-1] != '(' {
                pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] 
            }
            s = s[:len(s)-1]
        case specials[r] > 0: 
            for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
                pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
            }
            s = append(s, r)
        default:
          pofix = append(pofix, r)
        }
    }

    for len(s) > 0 {
        pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
    }

    return string(pofix)
}

// Regex match function from rega.go
// Match function on a string and postfix regexp
type state struct {
    symbol rune
    edge1 *state 
    edge2 *state
}

type nfa struct {
    initial *state
    accept *state
}

func poregtonfa(pofix string) *nfa {
    nfastack := []*nfa{}

    for _, r := range pofix {
        switch r {
        case '.':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            frag1.accept.edge1 = frag2.initial

            nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
        case '|':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            initial := state{edge1: frag1.initial, edge2: frag2.initial}
            accept := state{}
            frag1.accept.edge1 = &accept
            frag2.accept.edge1 = &accept

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        case '*':
            frag := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            accept := state{}
            initial := state{edge1: frag.initial, edge2: &accept}
            frag.accept.edge1 = frag.initial
            frag.accept.edge2 = &accept

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        default:     
            accept := state{}
            initial := state{symbol: r, edge1: &accept}

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        }
    }

    if len(nfastack) != 1 {
        fmt.Println("uh oh:", len(nfastack), nfastack)
    }
    return nfastack[0]
}

func addState(l []*state, s *state, a *state) []*state {
    l = append(l, s)

    if s != a && s.symbol == 0 {
        l = addState(l, s.edge1, a)
        if s.edge2 != nil {
            l = addState(l, s.edge2, a)
        }
    }
    return l
}

func pomatch(po string, s string) bool {
    
    ismatch := false

    ponfa := poregtonfa(po)
   
    current := []*state{}
    next := []*state{}

    current = addState(current[:], ponfa.initial, ponfa.accept)

    for _, r := range s {
        for _, c := range current {
            if c.symbol == r {
                next = addState(next[:], c.edge1, ponfa.accept)
            }
        }
        current, next = next, []*state{}
    }

    for _, c := range current {
        if c == ponfa.accept {
            ismatch = true
            break
        }
    }

    return ismatch
}

func main(){


    fmt.Println("test output")

    var swoption int
    var stringInput string
    var regexpInput string

    fmt.Println("Choose Between Infix and Postfix conversion \n Press 1: Infix \n Press 2: Postfix \n Press 3: Exit")
    fmt.Scan(&swoption)
    switch swoption {
        case 1:
        //case for infix conversion
        fmt.Println("case 1")

        fmt.Println("Enter the String and Infix Regular Expression to test")
        fmt.Scan(&stringInput, &regexpInput)

    case 2:
        //case for postfix conversion
        fmt.Println("case 2")

        fmt.Println("Enter the String and Postfix Regular Expression to test")
        fmt.Scan(&stringInput, &regexpInput)

    default:
        //default
    }
    //===============================================
    // Answer: ab.c
  /*  fmt.Println("infix: ", "a.b.c*")
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

    nfa := poregtonfa("ab.c*|")
    fmt.Println(nfa)

     fmt.Println(pomatch("ab.c|", "ab")) */
     //===============================================
}
