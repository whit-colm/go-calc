package nil

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/whitman-colm/go-lib/src/input"
	s "github.com/whitman-colm/go-lib/src/other"
	"math"
	"os"
)

////////////////////////////////////////
//Cylinder, cone and sphere calculator//
////////////////////////////////////////

var Version string = "\033[1;35m1.0.0\033[0m"

//Now with 53% more exportability!

//

func Startup() {
	//the function that's the "main menu" of sorts
	isDone := false
	doWhat := ""
	for isDone == false {
		//for loop so you don't have to come back here each time.
		//and a switch statment to force a responce.
		switch doWhat {
		case "a":
			fmt.Println("PROVIDE FUNCTION")
			doWhat = ""
		case "b":
			fmt.Println("PROVIDE FUNCTION")
			doWhat = ""
		case "d":
			isDone = true
		case "c":
			fmt.Println("PROVIDE FUNCTION")
			doWhat = ""
		default:
			fmt.Println(c.CL + c.B01 + "Quadratics calculator: " + Version)
			s.Spacer(2)
			fmt.Println(c.B1 + "What do you need solved?" + c.X)
			fmt.Println(c.Y + "Enter one of the letters in red:")
			s.Spacer(1)
			fmt.Println(c.R + "{A}" + c.B0 + " [OPTION A]")
			fmt.Println(c.R + "{B}" + c.B0 + " [OPTION B]")
			fmt.Println(c.R + "{C}" + c.B0 + " [OPTION C]")
			fmt.Println(c.R + "{D}" + c.B0 + " Exit system")
			doWhat = i.StringInput(c.M + ">>> " + c.B)

		}
	}
}

//Quadratic Formula: x=(-b ± sqrt(b^2-4ac))/2a
//Discriminant: b^2-4ac

func FindDiscriminant() {
	//This the the part where we declare A&B
	fmt.Println(c.CL + c.B1 + "What is a?")
	a := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is b?")
	b := i.DoubleInput(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is c?")
	cNumber := i.DoubleInput(c.M + ">>> " + c.B)
	//\\
	discAns := ((b * b) - (4 * a * cNumber))
	bb := b * b
	ac := a * cNumber
	ac4 := ac * 4
	//\\
	fmt.Printf("%s%s(%s%f%s)^2-4*%s%f%s*%s%f%s\n", c.CL, c.B01, c.B1, b, c.B01, c.B1, a, c.B01, c.B1, cNumber, c.B01)
	fmt.Printf("%s%f%s-4*%s%f%s*%s%f%s\n", c.B1, bb, c.B01, c.B1, a, c.B01, c.B1, cNumber, c.B01)
	fmt.Printf("%s%f%s-4*%s%f%s\n", c.B1, bb, c.B01, c.B1, ac, c.B01)
	fmt.Printf("%s%f%s-%s%f%s\n", c.B1, bb, c.B01, c.B1, ac4, c.B01)
	s.Spacer(1)
	fmt.Printf("%sThe discriminate is %f. ", c.O, discAns)
	solveQuad := 1
	if discAns > 0 {
		fmt.Printf("Meaning there are two solutions.\n")
	} else if discAns == 0 {
		fmt.Printf("Meaning there is one solution.\n")
	} else if discAns < 0 {
		fmt.Printf("Meaning there are no solutions.\n")
		solveQuad = 0
	} else {
		fmt.Println("ERROR. NOPE. BYE!")
		os.Exit(-1)
	}
	s.Go(1)
	//////////////////////////
	//Now to solve quadratic//
	//////////////////////////

	//Quadratic Formula: x=(-b ± sqrt(b^2-4ac))/2a

	if solveQuad == 1 {
		desc := math.Sqrt(discAns)
		negB := (-1 * b)
		firstTop := (negB + desc)
		secondTop := (negB - desc)
		quadA := (firstTop / (2 * a))
		quadB := (secondTop / (2 * a))
		a2 := (2 * a)
		//Show work
		fmt.Printf("%s%sx=(%s%f%s±%s%f%s)/2(%s%f%s)\n", c.CL, c.B01, c.B1, negB, c.B01, c.B1, desc, c.B01, c.B1, a, c.B01)
		fmt.Printf("%sx=(%s%f%s±%s%f%s)/%s%f%s\n", c.B01, c.B1, negB, c.B01, c.B1, desc, c.B01, c.B1, a2, c.B01)
		s.Spacer(1)
		fmt.Printf("%sx=(%s%f%s)/%s%f%s\n", c.B01, c.B1, firstTop, c.B01, c.B1, a2, c.B01)
		fmt.Printf("%sx=(%s%f%s)/%s%f%s\n", c.B01, c.B1, secondTop, c.B01, c.B1, a2, c.B01)
		s.Spacer(2)
		//Show answer
		if quadB == quadA {
			fmt.Printf("%sThe solution to this quadratic is %f\n", c.O, quadA)
		} else {
			fmt.Printf("%sThe solutions to this quadratic are %f and %f\n", c.O, quadA, quadB)
		}
		s.Go(1)
	}
}
