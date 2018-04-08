package main

import(
    "fmt"
    "bufio"
    "os"
    "strings"
)

//======================= intopost function from shunt.go ====================
//===================converts infix notation to postfix notation.==================
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

//===================== Regex match function from rega.go =======================
// =================Match function on a string and postfix regexp=================

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
func userInput() (string, error){
    reader := bufio.NewReader(os.Stdin)
    str, err := reader.ReadString('\n')

    return strings.TrimSpace(str),err
}


func main(){

    var swoption int
 
    fmt.Println("Choose Between Infix and Postfix conversion \n Press 1: Infix \n Press 2: Postfix \n Press 3: Exit")
    fmt.Scanln(&swoption)
    switch swoption {
        case 1:
        //case for infix conversion
        fmt.Print("Enter infix expression: ")
        readInfix, err := userInput()

        if err != nil{
            return
        }
       
		fmt.Println("Infix Expression entered: ", readInfix + "\nPostfix Expression: ", intopost(readInfix))

        postfix := intopost(readInfix)

		fmt.Println("Enter String to Match " + intopost(readInfix) + " Against: ")
		readStr, err := userInput()
		fmt.Println(pomatch(postfix,readStr))	

    case 2:
        //case for postfix conversion
        fmt.Println("Enter Postfix Expression: ")
        readPostfix, err := userInput()
      
        if err != nil{
              fmt.Println("uh oh:")
            return
        }

        fmt.Println("Postfix Expression entered: ", readPostfix)

        fmt.Println("Enter String to Match " + readPostfix + " Against: ")
        readStr, err := userInput()
		fmt.Println(pomatch(readPostfix,readStr))	

    default:
         fmt.Print("Goodbye!")
        //default
    }


}

