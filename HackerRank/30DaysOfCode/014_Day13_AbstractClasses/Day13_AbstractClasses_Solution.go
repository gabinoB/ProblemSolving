package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title  string
	Author string
}

type MyBook struct {
	*Book
	Price float64
}

func (b MyBook) display() {
	fmt.Print("Title: " + b.Title)
	fmt.Print("Author: " + b.Author)
	fmt.Print("Price: ", b.Price)
}

func main() {

	in := bufio.NewReader(os.Stdin)
	title, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	author, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	price, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	price = strings.TrimSpace(price) //Quita el espacio al final de la string para poder convertirlo correctamente a float.
	priceConverted, _ := strconv.ParseFloat(price, 8)

	b := Book{
		title,
		author,
	}

	mb := MyBook{
		&b,
		priceConverted,
	}

	mb.display()
}
