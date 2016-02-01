package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "input"
	"math"
	"strings"
	s "other"
)

///////////////////////////////////
//pythagorean theorem calculator://
////give it a & b, it finds c,/////
////give it b & c, it finds a!/////
///////////////////////////////////

var Version string = "\033[1;35m1.0.0\033[0m"
//Now with 53% more exportability!

func Startup() {
	//the function that's the "main menu" of sorts
	isDone := false
	for isDone == false{
		//for loop so you don't have to come back here each time.
		fmt.Println(c.B01 + "pythagorean theorem calculator " + Version)
		s.Spacer(2)
		fmt.Println(c.B1+"What do you need solved?"+c.X)
		fmt.Println(c.Y+"Enter one of the letters in red:")
		s.Spacer(1)
		fmt.Println(c.R+"{A}"+c.B0+"  I have A & B, I need C!")
		fmt.Println(c.R+"{B}"+c.B0+"  I have B & C, I need A!")
		fmt.Println(c.R+"{C}"+c.B0+"  Nothing! Stop talking to me perv!")
		doWhat := i.ReturnString(c.B+">>> "c.M)
	}
	//adds two spaces
}
