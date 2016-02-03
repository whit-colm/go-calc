package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-lib/src/input"
	s "github.com/whitman-colm/go-lib/src/other"
	"math"
)

////////////////////////////////////
//Distance and midpoint calculator//
////give it a & b, it finds c,//////
////give it b & c, it finds a!//////
////////////////////////////////////

var Version string = "\033[1;35m1.0.0\033[0m"

//Now with 53% more exportability!

func Startup() {
	//the function that's the "main menu" of sorts
	isDone := false
	doWhat := ""
	for isDone == false {
		//for loop so you don't have to come back here each time.
		//and a switch statment to force a responce.
		switch doWhat {
		case "a":
			findC()
			doWhat = ""
		case "b":
			findAorB()
			doWhat = ""
		case "c":
			isDone = true
		default:
			fmt.Println(c.CL + c.B01 + "distance and midpoint calculator " + Version)
			s.Spacer(2)
			fmt.Println(c.B1 + "What do you need solved?" + c.X)
			fmt.Println(c.Y + "Enter one of the letters in red:")
			s.Spacer(1)
			fmt.Println(c.R + "{A}" + c.B0 + "  Find distance")
			fmt.Println(c.R + "{B}" + c.B0 + "  Find midpoint")
			fmt.Println(c.R + "{C}" + c.B0 + "  Nothing! Stop talking to me perv!")
			doWhat = i.StringInput(c.M + ">>> " + c.B)
		}
	}
}

func findC() {
	//This the the part where we declare A&B
	fmt.Println(c.CL + c.B1 + "What is point x1?")
	x1 := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is point y1?")
	y1 := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is point x2?")
	x2 := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is point y2?")
	y2 := i.DoubleInput(c.M + ">>> " + c.B)
	//Now we math. c1 is the answer and all the others are just to show work
	x := x2 - x1
	y := y2 - y1
	xSq := x * x
	ySq := y * y
	c2 := xSq + ySq
	c1 := math.Sqrt(c2)
	//Now we show work
	fmt.Printf("%s%sThe formula is: sqrt((x2-x1)^2+(y2-y1)^2)\n", c.CL, c.B01)
	fmt.Printf("%sLet's do this:\nsqrt((%s%f%s-%s%f%s)^2-(%s%f%s-%s%f%s)^2\n",
		c.B01, c.B1, x2, c.B01, c.B1, x1, c.B01, c.B1, y2, c.B01, c.B1, y1, c.B01)
	fmt.Printf("%ssqrt((%s%f%s)^2+(%s%f%s)^2)\n", c.B01, c.B1, x, c.B01, c.B1, y, c.B01)
	fmt.Printf("%ssqrt(%s%f%s+%s%f%s)\n", c.B01, c.B1, xSq, c.B01, c.B1, ySq, c.B01)
	fmt.Printf("%ssqrt(%s%f%s)\n", c.B01, c.B1, c2, c.B01)
	s.Spacer(1)
	fmt.Printf("%sSquared distance is %f\n", c.O, c1)
	fmt.Printf("%sUnsquared distance is %f\n", c.O, c2)
	s.Go(1)
}

func findAorB() {
	//This the the part where we declare A&B
	fmt.Println(c.CL + c.B1 + "What is point x1?")
	x1 := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is point y1?")
	y1 := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is point x2?")
	x2 := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is point y2?")
	y2 := i.DoubleInput(c.M + ">>> " + c.B)
	//Now we math. c1 is the answer and all the others are just to show work
	x := x2 + x1
	y := y2 + y1
	xSq := x / 2
	ySq := y / 2
	//Now we show work
	fmt.Printf("%s%sThe formula is: ((x2+x1)/2),((y2+y1)/2)\n", c.CL, c.B01)
	fmt.Printf("%sLet's do this:\n((%s%f%s+%s%f%s)/2),((%s%f%s+%s%f%s)/2)\n",
		c.B01, c.B1, x2, c.B01, c.B1, x1, c.B01, c.B1, y2, c.B01, c.B1, y1, c.B01)
	fmt.Printf("%s((%s%f%s)/2),(%s%f%s)/2)\n", c.B01, c.B1, x, c.B01, c.B1, y, c.B01)
	fmt.Printf("%s(%s%f%s,%s%f%s)\n", c.B01, c.B1, xSq, c.B01, c.B1, ySq, c.B01)
	s.Spacer(1)
	fmt.Printf("%sthe midpoint is (%f,%f).\n", c.O, xSq, ySq)
	s.Go(1)
}
