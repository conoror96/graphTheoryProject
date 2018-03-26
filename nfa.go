//Thompson's Contruction
//This is an algorithm for transforming a regular expression into an equivalent NFA
//Ref: https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//to do: change to match regular expressions against a string, Comment code

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


    return nfastack[0]
}

func main() {
    nfa := poregtonfa("ab.c*|")
    fmt.Println(nfa)
}