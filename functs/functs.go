package main

import (
	"fmt"
	"os"
)

/*
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.
s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt
the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo,
and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given
values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t,
representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.

*/

// GenDisplaceFn does some fun stuff with taking physics inputs and outputs a function
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
	return fn
}

// PrintError prints errors
func PrintError(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main() {

	var a, v0, s0, t float64
	var err error

	fmt.Print("Enter acceleration (m/s^2):")
	_, err = fmt.Scan(&a)
	PrintError(err)

	fmt.Print("Enter initial velocity (m/s):")
	_, err = fmt.Scan(&v0)
	PrintError(err)

	fmt.Print("Initial Position (m):")
	_, err = fmt.Scan(&s0)
	PrintError(err)

	disFunct := GenDisplaceFn(a, v0, s0)

	fmt.Print("\nPlease enter a time (s):")
	_, err = fmt.Scan(&t)

	fmt.Printf("Displacement %f m\n", disFunct(t))

}
