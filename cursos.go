package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// Si una funcion tiene un receptor de tipo puntero hay que poner todas las funciones de tipo puntero
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c *Course) PrintClasses() { // indicando que es un método vinculado a la estructura Course
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
