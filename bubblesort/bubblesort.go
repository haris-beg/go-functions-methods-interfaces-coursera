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

func main() {
    fmt.Print("Enter a sequence of up to 10 integers, separated by spaces: ")
    r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

    // split the line into separate components
    fields := strings.Fields(line)
    if len(fields) > 10 {
        log.Fatal("Too many arguments")
    }

    sliceOfInts := make([]int, len(fields))

    for i, f := range fields {
		sliceOfInts[i], err = strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
	}

    fmt.Println("Before sort: ", sliceOfInts)
    BubbleSort(sliceOfInts)
    fmt.Println("After sort: ", sliceOfInts)

}

func BubbleSort(sliceOfInts []int) {
    // Perform Bubble Sort on the slice of integers.
    // Repeatedly Swap adjacent elements if they are not in order.
    // The algorithm needs one whole pass without any swap to know that sorting is complete
    for {
        numOfSwaps := 0
        for i := 0; i < len(sliceOfInts) - 1 ; i++ {
            if sliceOfInts[i] > sliceOfInts[i+1] {
                Swap(sliceOfInts, i)
                numOfSwaps++
            }
        }
        if numOfSwaps == 0 {
            break
        }
    }
}

func Swap(sliceOfInts []int, i int) {
    temp := sliceOfInts[i]
    sliceOfInts[i] = sliceOfInts[i+1]
    sliceOfInts[i+1] = temp
}