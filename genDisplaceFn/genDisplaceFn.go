/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take 
one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    // request user input
    fmt.Println("Enter the acceleration, initial velocity, and initial displacement, separated by space: ")
    
    // read one line from standard input
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    exitOnError(err)

    // split the line on whitespace into separate tokens
    fields := strings.Fields(line)
    if len(fields) != 3 {
        log.Fatal("Invalid number of input elements. Exiting.")
    }

    // convert the tokensinto float64 variables
    acceleration, err := strconv.ParseFloat(fields[0], 64)
    exitOnError(err)
    initVelocity, err := strconv.ParseFloat(fields[1], 64)
    exitOnError(err)
    initDisplacement, err := strconv.ParseFloat(fields[2], 64)
    exitOnError(err)

    // generate the function
    fn := GenDisplaceFn(acceleration, initVelocity, initDisplacement)
    
    fmt.Print("Enter time in seconds: ")
    var time float64
    n, err := fmt.Scanf("%f", &time)
    exitOnError(err)
    if n != 1 {
        log.Fatal("Invalid input")
    }

    // use the function to calculate the displacement
    displacement := fn(time)
    fmt.Println("Displacement after ",time, " seconds = ", displacement)
}

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
    return func(t float64) float64 {
        return (0.5 * a * t * t) + (v0 * t) + s0
    }
}

func exitOnError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}