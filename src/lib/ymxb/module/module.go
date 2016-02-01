package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	u "github.com/whitman-colm/go-1/utils/other"
	"strconv"
)

////////////////
//EXPLANATION!//
////////////////
/*
*This is so clunky because I made this in late 2015 and have
*no intention at this time of rewriting this with my new
*knowledge of golang. I might redo all of this when I have
*more time but for now I'm going to let it do it's thing. soz.
 */

func PointSlope() {
	fmt.Print(c.CL)
	strx1, err := i.Prompt(c.G + "Enter x1 point\n" + c.B + ">" + c.M)
	u.QuitAtError(err)
	stry1, err := i.Prompt(c.G + "Enter y1 point\n" + c.B + ">" + c.M)
	u.QuitAtError(err)
	strm, err := i.Prompt(c.G + "Enter slope\n" + c.B + ">" + c.M)
	u.QuitAtError(err)
	//all inputs now provided, to be converted.
	x1, err := strconv.ParseFloat(strx1, 64)
	u.QuitAtError(err)
	y1, err := strconv.ParseFloat(stry1, 64)
	u.QuitAtError(err)
	m, err := strconv.ParseFloat(strm, 64)
	u.QuitAtError(err)
	//all data points provided and converted, now to math-a-tise.
	fmt.Println(c.CL, c.G+"The formula is")
	fmt.Println("y = m(x) + b")
	u.Spacer(3)
	//to find "b"
	fmt.Println(y1, " = ", m, "(", x1, ") + b")
	filler := x1 * m
	//multiplies x1 and m to filler.
	u.Spacer(1)
	fmt.Println(y1, " = ", filler, " + b")
	fmt.Println("-", filler, "---------------|")
	//Shows subtraction
	u.Spacer(1)
	b := y1 - filler
	fmt.Println(c.B2+"b =  ", b)
	u.Spacer(3)
	fmt.Println(c.B3+"y = ", m, "(x) + ", b)
	u.Go(1)
	//prints out completed statment, ends function
}

func Standard() {
	fmt.Println(c.CL, c.G, "Sorry m9, this hasn't been devopled yet...")
	u.Go(1)
}

func TwoPoints() {
	fmt.Print(c.CL)
	strx1, err := i.Prompt(c.G + "Enter x1 point\n" + c.B + ">" + c.M)
	//assigns a command-line input to strx1.
	//i.Prompt returns a string "strx1" and an error "err"
	u.QuitAtError(err)
	//Tests if there is an error in the input, and halts the program before it crashes
	stry1, err := i.Prompt(c.G + "Enter y1 point\n" + c.B + ">" + c.M)
	u.QuitAtError(err)
	strx2, err := i.Prompt(c.G + "Enter x2 point\n" + c.B + ">" + c.M)
	u.QuitAtError(err)
	stry2, err := i.Prompt(c.G + "Enter y2 point\n" + c.B + ">" + c.M)
	u.QuitAtError(err)
	//provides the data values. will be converted next 4 lines. "i" pulls from my input function
	//variable err will be processed on-site w/ the input function next version.
	x1, err := strconv.ParseFloat(strx1, 64)
	//Converts strx1 in to a 64-bit float.
	u.QuitAtError(err)
	y1, err := strconv.ParseFloat(stry1, 64)
	u.QuitAtError(err)
	x2, err := strconv.ParseFloat(strx2, 64)
	u.QuitAtError(err)
	y2, err := strconv.ParseFloat(stry2, 64)
	u.QuitAtError(err)
	/*
	*Now for something to actually happen! all variables have been declared!
	 */
	fmt.Println(c.CL + c.G + "The Formula is:" + c.R)
	fmt.Println(c.B1, "y2 - y1")
	fmt.Println(c.B1, "-------")
	fmt.Println(c.B1, "x2 - x1")

	u.Spacer(3)
	//adds three lines blank diffrence.
	fmt.Println(c.R, y2, "-", y1)
	fmt.Println(c.G + "-----------")
	fmt.Println(c.R, x2, "-", x1)
	//Prints out an incomplete equation
	u.Spacer(2)
	num := y2 - y1
	den := x2 - x1
	//the numerator and denominator of the equation.
	fmt.Print(c.R)
	fmt.Println(num)
	fmt.Println(c.G + "-----")
	fmt.Print(c.R)
	fmt.Println(den)
	u.Spacer(2)
	if den == 0 {
		fmt.Println(c.B2, "The slope is undefined")
	} else {
		m := num / den
		fmt.Println(c.B3, "The slope is ", m, ".")
		u.Go(1)
		/////////////////////////////////////
		//adds a break and clears the page.//
		/////////////////////////////////////
		fmt.Println(c.G + "Our y-intercept equation is:")
		fmt.Println(c.B1, "y = m(x) + b"+c.G)

		//all data points provided and converted, now to math-a-tise
		u.Spacer(3)
		//to find "b"
		fmt.Println(y1, " = ", m, "(", x1, ") + b")
		filler := x1 * m
		//multiplies x1 and m to filler
		fmt.Println(y1, " = ", filler, " + b")
		fmt.Println("-(", filler, ")---------------|")
		//Shows subtraction
		u.Spacer(1)
		b := y1 - filler
		fmt.Println(c.B2, "b = ", b)
		u.Spacer(3)
		fmt.Println(c.B3+"y = ", m, "(x) + ", b)
		u.Go(1)
		//prints out completed statment, ends function
	}
}
