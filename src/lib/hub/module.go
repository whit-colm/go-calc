package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-lib/src/input"
	s "github.com/whitman-colm/go-lib/src/other"
	ymxb "lib/ymxb"
)

//////////////////////////////////////////////////
//This is just a hub for each of the modules//////
//The real programming code is in another castle//
//////////////////////////////////////////////////

func Start() {
	//the only function in this
	isDone := false
	doWhatNow := ""
	for isDone == false {
		//So that you don't have to restart the program each time
		switch doWhatNow {
		case "a":
			ymxb.Function()
		case "b":
			//call another thing
		case "c":
			isDone := true
		default:
			//forces a responce
			fmt.Println(c.CL + c.B00 + "Mathmatic solver program thing.")
			s.Spacer(2)
			//Adds two lines
			fmt.Println(c.B1 + "What are you too lazy to do today?")
			fmt.Println(c.Y + "Select one of the letters in red")
			s.Spacer(1)
			fmt.Println(c.R + "{A}" + c.B0 + "Find a linear slope equation")
			fmt.Println(c.R + "{B}" + c.B0 + "Find the hypotonuse of a right triangle")
			fmt.Println(c.R + "{C}" + c.B0 + "Nothing go away")
			s.Spacer(1)
			doWhatNow := i.StringInput(c.M + ">>>" + c.B)
		}
	}
}
