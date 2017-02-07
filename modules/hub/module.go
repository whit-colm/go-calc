package nil

import (
	"fmt"

	dist "../disAndMid"
	quad "../quadratic"
	tria "../triangle"
	ymxb "../ymxb"
	m "github.com/marcmak/calc/calc/solver.go"
	c "github.com/skilstak/go-colors"
	i "github.com/whitman-colm/go-lib/src/input"
	s "github.com/whitman-colm/go-lib/src/other"
	"strconv"
)

// import
//ccs "lib/ccs"
//"strings"

//////////////////////////////////////////////////
//This is just a hub for each of the modules//////
//The real programming code is in another castle//
//////////////////////////////////////////////////

//Also it now does basic math from command line

func Start() {
	solve := ""
	fmt.Println(c.CL + "Type " + c.O + "fn" + c.X + " for math functions and " + c.O + "q" + c.X + " to quit.")

	for solve != "q" {
		solve = i.StringInput("~> " + c.B01)
		if solve == "fn" {
			mathFuncs()
		} else if solve == "q" {
			fmt.Println(c.X + "Exiting calculator. Have a great day.")
		} else if m.ContainsLetter(solve) == false {
			ans := m.Solve(solve)
			ansFmt := strconv.FormatFloat(ans, 'E', -1, 64)
			fmt.Println(c.B01 + " [ " + c.G + ansFmt + c.B01 + " ]")
		} else {
			fmt.Println(c.R + "Input could not be read by system. Please try again.")
		}
	}
}
func mathFuncs() {
	//the only function in this
	isDone := false
	doWhatNow := ""
	for isDone == false {
		//So that you don't have to restart the program each time
		switch doWhatNow {
		case "a":
			ymxb.Function()
			//Calls the ymxb library's startup function. see module for more details.
			doWhatNow = ""
			//We have to reset doWhatNow lest we want the same module on repeat
		case "b":
			tria.Startup()
			doWhatNow = ""
		case "c":
			dist.Startup()
			doWhatNow = ""
		case "d":
			// ccs.Startup()  // not found?
			doWhatNow = ""
		case "e":
			quad.FindDiscriminant()
			doWhatNow = ""
		case "f":
			isDone = true
		default:
			//forces a responce
			fmt.Println(c.CL + c.X + "Functions")
			fmt.Println(c.Y + "Select one of the letters in red")
			s.Spacer(1)
			//adds a line
			fmt.Println(c.R + "{A}" + c.B0 + "  Find a linear slope equation")
			fmt.Println(c.R + "{B}" + c.B0 + "  Find the hypotonuse of a right triangle")
			fmt.Println(c.R + "{C}" + c.B0 + "  Find distance or midpoint")
			fmt.Println(c.R + "{D}" + c.B0 + "  Find volume of cone, cylinder or sphere")
			fmt.Println(c.R + "{E}" + c.B0 + "  Quadratics")
			fmt.Println(c.R + "{F}" + c.B0 + "  Nothing go away")
			s.Spacer(1)
			doWhatNow = i.StringInput(c.B + ">>>" + c.X)
		}
	}
}
