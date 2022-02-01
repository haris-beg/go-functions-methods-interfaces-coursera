package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user,
prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:
food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion, and
the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.
*/

type Animal struct {
	food, locomotion, noise string
}

func (animal Animal) Eat()  {
	fmt.Println(animal.food)
}

func (animal Animal) Move()  {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak()  {
	fmt.Println(animal.noise)
}

var (
	cow Animal = Animal{
		food: "grass", 
		locomotion: "walk",
		noise: "moo",
	}
	
	bird Animal = Animal{
		food: "worms", 
		locomotion: "fly",
		noise: "peep",
	}
	
	snake Animal = Animal{
		food: "mice", 
		locomotion: "slither",
		noise: "hsss",
	}
)

func main() {
	// create an infinite loop 
	for {
		fmt.Print("> ")
		line := readStdinLine()
		fields := splitStringIntoTokens(line, 2)
		respondToInput(fields)
	}
}

func respondToInput(fields []string) {
	switch fields[0] {
	case "cow":
		checkVerb(cow, fields[1])
	case "bird":
		checkVerb(bird, fields[1])
	case "snake":
		checkVerb(snake, fields[1])
	default:
		log.Fatal("Invalid animal entry!")
	}
}

func checkVerb(animal Animal, verb string) {
	switch verb {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		log.Fatal("Invalid action entered!")
	}
}

func splitStringIntoTokens(line string, desiredNum int) []string {
	// split line on whitespace into tokens
	fields := strings.Fields(line)

	// check for correct number of tokens
	if len(fields) != desiredNum {
		log.Fatal("Invalid number of input arguments!")
	}

	return fields
}

func readStdinLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
        log.Fatal(err)
    }
	return line
}