package lab1

import (
	"testing"
	. "gopkg.in/check.v1"
	"fmt"
)
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})


func (s * TestSuite) TestConversion(c *C) {

	expressions := [5]string{"4 4 +",
                                 "5 8 * 9 3 / - 6 3 ^ - 3 / 3 7 8 * - 3 ^ +",
                                 "  3     1     /   ",
                                 "4 4 + 7 *",
                                 "5"}
							 
	outcome := [5]string{"4 + 4",
                             "((((5 * 8) - (9 / 3)) - (6 ^ 3)) / 3) + ((3 - (7 * 8)) ^ 3)",
                             "3 / 1",
                             "(4 + 4) * 7",
                             "5"}
						 
	for i := 0; i < len(expressions); i++ {
		res, err := PostfixToInfix(expressions[i])
		if c.Check(err, IsNil) {
			c.Assert(res, Equals, outcome[i])
			c.Assert(res, NotNil)
		}
	}


	false_expressions := [9]string{"",
				       "      ",
                                       "6+",
                                       "q",
                                       "4d",
                                       "5 6+",
                                       "+ 8",
                                       "4 4 + + 4",
                                       "4 5 + 5 &"}
							 
	for i := 0; i < len(false_expressions); i++ {
		_, err := PostfixToInfix(false_expressions[i])
		if c.Check(err, NotNil) {
			c.Assert(err, ErrorMatches, "Incorrect expression")
		}
	}
	
}

func ExamplePostfixToInfix() {

	res, err := PostfixToInfix("5 1 / 7 3 * + 2 ^ 4 ^")
	if err == nil {
		fmt.Println(res)
	}

	// Output:
	// ((5 / 1) + (7 * 3)) ^ 2 ^ 4
}