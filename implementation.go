package lab1
 
import (
    "strings"
    "regexp"
    "errors"
)


func PostfixToInfix(text string) (string, error) {
	regNum := regexp.MustCompile(`\s*[0-9]+\s*`)
        regSymb := regexp.MustCompile(`\s*[+-/*^]+\s*`)
	matched, _ := regexp.Match(`[0-9][+-/*^]|[+-/*^][0-9]|[+-/*^]{2,}|[^0-9+-/*^\s+]`, []byte(text))
	if(matched ||len(regNum.FindAll([]byte(text), -1)) != len(regSymb.FindAll([]byte(text), -1)) + 1){
		return "", errors.New("Incorrect expression")
    }

    var stack []string
	var s1 string
	var s2 string
	var temp string
	numCount := 0
	opCount := 0
    for _, tok := range strings.Fields(text) {
		if (tok == "*" || tok == "/" || tok == "^" || tok == "+" || tok == "-") {
			opCount++
			if (numCount > opCount) {
				s1, stack = stack[len(stack)-1], stack[:len(stack)-1]
				s2, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if tok == strings.Fields(text)[len(strings.Fields(text)) - 1] {
					temp = s2 + " " + tok + " " + s1
				} else {
					temp = "(" + s2 + " " + tok + " " + s1 + ")"
				}
				stack = append(stack, temp)
			} else {
				return "", errors.New("Incorrect expression")
			}
		} else {
			numCount++
			stack = append(stack, tok)
		}
    }
	return stack[0], nil
}
