//read a match function on a string
//Ref: https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
//to make work with infix regexp change first step of match, convert infix regexp to postfix regexp
// add + and * operator

package main

import (
    "fmt"
)

//struct to store states and arrows
//max arrows from every state is 2
type state struct {
    symbol rune
    edge1 *state //arrow points at state
    edge2 *state
}

//build a list of state structs that represent nfa
type nfa struct {
    initial *state
    accept *state
}


func poregtonfa(pofix string) *nfa {
    nfastack := []*nfa{}

    //loop through postfix regexp. 
    for _, r := range pofix {
        switch r {
        case '.':
            //pop 2 things off nfa nfastack
            frag2 := nfastack[len(nfastack)-1]
            //get rid of last thing off the nfastack
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            frag1.accept.edge1 = frag2.initial

            //& = address of that instance, push a new thing to the stack
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

    // if any state has 0 val of rune, then this state has e arrows coming from it
    if s != a && s.symbol == 0 {
        l = addState(l, s.edge1, a)
        if s.edge2 != nil {
            l = addState(l, s.edge2, a)
        }
    }

    return l
}

//takes postfixregexp and a string (s) 
//
func pomatch(po string, s string) bool {
    //false by default
    ismatch := false

    //
    ponfa := poregtonfa(po)
    //create array of pointers to state
    current := []*state{}
    //keeps track of all states s can get to next
    next := []*state{}

    //create new function add state. pass current to it
    //current is list of states when we start off
    current = addState(current[:], ponfa.initial, ponfa.accept)

    //loop through string. then put all current states into next from current and clear next
    for _, r := range s {
        for _, c := range current {
            if c.symbol == r {
                next = addState(next[:], c.edge1, ponfa.accept)
            }
        }
        current, next = next, []*state{}
    }

    //loop through current array. 
    for _, c := range current {
        if c == ponfa.accept {
            ismatch = true
            break
        }
    }

    return ismatch
}

func main() {
    fmt.Println(pomatch("ab.c|", "kk"))
}