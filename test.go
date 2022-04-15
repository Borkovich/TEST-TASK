package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var name, color string
	var weight int

	fmt.Println("Введите название")
	fmt.Scanf("%s\n", &name)

	fmt.Println("Введите вес")
	fmt.Scanf("%d\n", &weight)
	fmt.Scanf("%s\n", &color)
	fmt.Println(chooseColor(color))

	if len(name) < 3 || len(name) > 120 {
		log.Fatal("name too long or short")
	} else {
		fmt.Println(name)
	}
	if weight < 1 || weight > 500 {
		log.Fatal("weight to high or small")
	} else {
		fmt.Println(weight)
	}

}
func chooseColor(color string) (string, error) {
	switch color {
	case "r":
		return "Красный", nil
	case "g":
		return "Зеленый", nil
	case "b":
		return "Синий", nil
	default:
		return "Вы не выбрали цвет!", errors.New("choose color")
	}

}
