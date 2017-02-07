package nil

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/whitman-colm/go-lib/src/input"
	u "github.com/whitman-colm/go-lib/src/other"
	s "./module"
)

////////////////
//EXPLANATION!//
////////////////
/*
* This is so clunky because I made this in late 2015 and have
* no intention at this time of rewriting this with my new
* knowledge of golang. I might redo all of this when I have
* more time but for now I'm going to let it do it's thing. soz.
* I am also aware of the fact I didn't document any of this.
 */

func Function() {
	var done bool = false

	fmt.Println(c.CL)
	fmt.Println(c.B3, "Welcome to the Liner Solver by Whitman Huntley.")
	fmt.Println(c.B3, "This calculator supports all rational numbers as float64 datatypes.")
	fmt.Println(c.B3, "This program is text based.To make anything function, enter a keyword and then press enter.")
	u.Spacer(1)
	fmt.Println(c.V, "ERROR CHECK:", c.B00, "Don't mind me, I'm just here to make sure the program \n doesn't break.")
	u.Go(1)

	for done == false {
		fmt.Println(c.CL + c.X + c.G + "What data do you have?")
		fmt.Println(c.R + "a" + c.G + ":  Two pairs of points. (x1, y1), (x2, y2)")
		fmt.Println(c.R + "b" + c.G + ":  The slope and a point. m, (x1, y1)")
		fmt.Println(c.R + "c" + c.G + ":  None, I'd like to leave.")
		mode, err := i.Prompt(c.G + "Select mode using the red key above\n" + c.B + ">" + c.M)
		u.QuitAtError(err)

		switch mode {
		//mode is defined as an input string above.
		case "a", "1":
			s.TwoPoints()
		case "b", "2":
			s.PointSlope()
		case "c", "3":
			fmt.Println(c.Y, "Okey! bye-bye!")
			done = true
		default:
			fmt.Println(c.B, "That statment isn't valid...")
			u.Go(1)
		}
	}
}
