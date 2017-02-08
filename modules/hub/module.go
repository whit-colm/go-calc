package nil

import (
	"fmt"

	m "../../calc"
	dist "../disAndMid"
	quad "../quadratic"
	tria "../triangle"
	ymxb "../ymxb"
	c "github.com/skilstak/go-colors"
	i "github.com/whitman-colm/go-lib/src/input"
	s "github.com/whitman-colm/go-lib/src/other"
	//m "markmac/calc"
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
	fmt.Println(c.CL + c.X + "Get help by using " + c.O + "?" + c.X + ".")

	for solve != "q" {
		solve = i.StringInput("~> " + c.B01)
		if solve == "fn" {
			mathFuncs()
			fmt.Println(c.CL + c.X + "Get help by using " + c.O + "?" + c.X + ".")
		} else if solve == "?" || solve == "help" || solve == "h" {
			help()
			fmt.Println(c.CL + c.X + "Get help by using " + c.O + "?" + c.X + ".")
		} else if solve == "q" {
			fmt.Println(c.X + "Exiting calculator. Have a great day.")
		} else if solve == "cl" || solve == "clear" || solve == "x" {
			fmt.Println(c.CL + c.X + "Get help by using " + c.O + "?" + c.X + ".")
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

func help() {
	fmt.Println(c.CL + c.O + "GO CALCULATE USAGE GUIDE AND COMMANDS")
	s.Spacer(1)
	fmt.Println(c.X + "go-calc has the ability to solve simple equasions directly from the prompt.\nSimply type in an expression and (hopefully) go-calc will spit out an answer.\nIt can also evaluate certian equasions. Commands are below:")
	s.Spacer(2)
	fmt.Println(c.B0 + "?" + c.X + ", " + c.B0 + "h" + c.X + ", " + c.B0 + "help   " + c.X + ": " + c.B01 + "Open help pane. (you should know this one because you're using it now boi-o.)")
	fmt.Println(c.B0 + "cl" + c.X + ", " + c.B0 + "x" + c.X + ", " + c.B0 + "clear " + c.X + ": " + c.B01 + "Clear screen.)")
	fmt.Println(c.B0 + "q           " + c.X + " : " + c.B01 + "Quit the program.")
	fmt.Println(c.B0 + "fn          " + c.X + " : " + c.B01 + "Open functions menu.")
	s.Go(1)
}
