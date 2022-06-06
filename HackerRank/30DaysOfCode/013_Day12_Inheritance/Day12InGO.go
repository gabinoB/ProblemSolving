package main

import (
	"fmt"
)

type Person struct {
	firstName      string
	lastName       string
	identification int32
}

type profesor struct {
	*Person
	testScores []int32
}

func (pp profesor) calculate() string {
	sum := 0
	result := 0
	for i := 0; i < len(pp.testScores); i++ {
		sum = sum + int(pp.testScores[i])
	}
	result = sum / len(pp.testScores)

	if result <= 100 && result >= 90 {
		return "O"
	} else if result < 90 && result >= 80 {
		return "E"
	} else if result < 80 && result >= 70 {
		return "A"
	} else if result < 70 && result >= 55 {
		return "P"
	} else if result < 55 && result >= 40 {
		return "D"
	} else {
		return "T"
	}
}

func (p Person) printPerson() {
	fmt.Printf("Name: %s, %s \nID: %v", p.lastName, p.firstName, p.identification)
}

func main() {

	var firstName string
	var lastName string
	var id int32
	fmt.Scan(&firstName, &lastName, &id)

	var amountOfTestScores uint8
	fmt.Scan(&amountOfTestScores)

	testScores := make([]int32, amountOfTestScores)

	for i := 0; i < int(amountOfTestScores); i++ {
		var v int32
		fmt.Scan(&v)
		testScores = append(testScores, v)
	}

	p := Person{
		firstName,
		lastName,
		id,
	}

	pp := profesor{
		&p,
		testScores,
	}

	p.printPerson()
	grade := pp.calculate()

	fmt.Printf("\nGrade: %s", grade)

}
