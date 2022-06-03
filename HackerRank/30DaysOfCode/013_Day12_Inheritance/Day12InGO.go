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
		//sum = sum + pp.testScores[i]
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

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(id)

	//scanner := bufio.NewScanner(os.Stdin)
	//
	//data := scanner.Text()
	//
	//fmt.Println(data)

	//var data string
	//
	//for i := 0; i < 3; i++ {
	//	data = data + scanner.Text()
	//}
	//
	//fmt.Println(data)

	//var data string
	//fmt.Scanln(&data)
	//
	//fmt.Println(data)
	//fmt.Println(data[0])
	//fmt.Println(data[1])
	//fmt.Println(data[2])

	//reader := bufio.NewReader(os.Stdin)
	//data, err := reader.ReadSlice('\n')
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(data)

	//var firstName string
	//var lastName string
	//var identification int32

	//in := bufio.NewReader(os.Stdin)
	//data, _ := in.ReadString('\n')

	//var data [3]string
	//fmt.Scan(&data)
	//firstName := string(data[0])
	//lastName := data[1]
	//identification := int32(data[2])
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Printf("%v %v %v", firstName, lastName, identification)

	//gabino := Person{
	//	firstName:      "Gabino",
	//	lastName:       "Ballay",
	//	identification: 37562213,
	//}
	//
	//var arr []int32 //Problemas con el indice y el tamaño de los arrays y slices
	//arr[0] = 1      //Ver bien el tema de los tipos de datos estaticos y dinamicos.
	//arr[1] = 2      //Necesito que el array sea dinamico.
	//
	//gabinoProfesor := profesor{
	//	&gabino,
	//	arr,
	//}
	//
	//gabino.printPerson()
	//fmt.Printf("%v", gabinoProfesor.testScores)

}
