package pkg

import "strings"

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push: add a new value onto the stack
func (s *Stack)Push(str string){
	*s = append(*s, str)
}

// Pop: remove and return top element of stack. Resturn false if stack is empty
func (s *Stack)Pop()(string, bool){
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// BalanceExpresion: check if the expression has the correct special characters when open and close
// the special characters are "(", "{", "[", ")", "}", "]"
func BalanceExpresion(expr string) bool {
	var checkExpr Stack

	for _, c := range strings.Split(expr, "") {
		if c == "(" || c == "[" || c == "{" {
			checkExpr.Push(c)
		} else {
			if c == ")" {
				if pop, _ := checkExpr.Pop(); pop != "("{
					return false
				}
			}
			if c == "]" {
				if pop, _ := checkExpr.Pop(); pop != "["{
					return false
				}
			}
			if c == "}" {
				if pop, _ := checkExpr.Pop(); pop != "{"{
					return false
				}
			}
		}
	}

	return checkExpr.IsEmpty()
}
