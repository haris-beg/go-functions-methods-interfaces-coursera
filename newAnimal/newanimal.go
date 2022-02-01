/*

Write a program which allows the user to create a set of animals and 
to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type, but the types of animals are restricted
to either cow, bird, or snake. 
The following table contains the three types of animals and their associated data.

Your program should present the user with a prompt, “>”, 
to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response,
and print out a new prompt on a new line. Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings.
The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal
and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings.
The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal,
either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
which take no arguments and return no values.
The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak()
so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var animals = make(map[string]Animal)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// create an infinite loop 
	for {
		fmt.Print("> ")
		line := readStdinLine()
		fields := splitStringIntoTokens(line, 3)
		respondToInput(fields)
	}
}

func respondToInput(fields []string) {
	command := fields[0]
	switch command {
	case "newanimal":
		addNewAnimal(fields[1], fields[2])
	case "query":
		queryAnimal(fields[1], fields[2])
	default:
		log.Fatal("Invalid command!")
	}
}

func queryAnimal(name, verb string) {
	switch verb {
	case "eat":
		animals[name].Eat()
	case "move":
		animals[name].Move()
	case "speak":
		animals[name].Speak()
	default:
		log.Fatal("Invalid verb: ", verb)
	}
}

func addNewAnimal(animalName, animalType string) {
	switch animalType {
	case "cow":
		animals[animalName] = Cow{}
	case "bird":
		animals[animalName] = Bird{}
	case "snake":
		animals[animalName] = Snake{}
	default:
		log.Fatal("Invalid animal type: ", animalType)
	}
	fmt.Println("Created it!")
}

func readStdinLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
        log.Fatal(err)
    }
	return line
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