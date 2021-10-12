package main

import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	Float()
	Floatd()
	b := []byte("hello")             // конвертация в байт
	bn := []byte{11, 12, 13, 14, 15} // массив байт
	fmt.Println(b)
	fmt.Println(bn)
	Iota()

	// руны и байты стринга
	// если строку разбираем посимвольно, это будет руна
	op := "string"
	for i, c := range op {
		// приводит руну к стрингу
		fmt.Println(string(c), c)
		fmt.Println(string(op[i]), c)
		fmt.Println(string(op[i]), unsafe.Sizeof(c), "byte")
	}

	// formated output
	// For development marshal unmarshal, api, html etc.
	// Don't use it in production
	s := `
	Hello
	my
	dear
	friend!
	`
	fmt.Println(s)

	var a float32
	fmt.Print("enter a: ")
	n, err := fmt.Scan(&a)
	if err == nil && n > 0 {
		// количество знаков после запятой
		fmt.Printf("a = %0.2f \n", a)
	} else {
		log.Println("Error!")
	}

	op = ""
	switch op {
	case "+", "*":
		fmt.Println("operator + or * op = ", op)
		fallthrough // пробросит ниже и будет напечатана инструкция ниже
	case "-":
		fmt.Println("operator -, op = ", op)
	default:
		fmt.Println("default")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for iterator, v := range []int{5, 4, 3, 2, 1} {
		fmt.Println(iterator, v)
	}

	// Бесконечный цикл
	i := 0
	for {
		fmt.Println(i + 1)
	}
}
