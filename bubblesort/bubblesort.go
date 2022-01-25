package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

const (
    format = "%d %d %d %d %d %d %d %d %d %d"
)

    // mySlice := make([]int, 10)

func main() {
    fmt.Print("Please enter a sequence of up to 10 integers: ")
    r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(line)

    // split the line into separate components
    fields := strings.Fields(line)
    if len(fields) > 10 {
        log.Fatal("Too many arguments")
    }

    ints := make([]int, len(fields))

    for i, f := range fields {
		ints[i], err = strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
	}

    fmt.Println(ints)

}