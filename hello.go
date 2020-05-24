package main

import (
	"fmt"
)

// func main() {
// 	testArrSlice()
// 	testMap()
// 	testTypes()
// }

func testArrSlice() {
	// array fijo de 3 posiciones con valores iniciales 1, 2, 3
	a := [3]int{1, 2, 3}

	// genero un slice usando a desde su posicion 0 hasta el final
	s := a[0:]

	// puedo agregarle un valor usando append (no muta el array original)
	s = append(s, 88)

	// si modifico con un slice el array, vemos que modifica el original
	s1 := a[:]
	s2 := a[:]
	s1[0] = 100 // s2[0] ahora tiene 100

	fmt.Println(s, s1, s2)

	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func testMap() {
	// un map es un hashtable -> key: value
	m := map[int]string{0: "Hola Mundo", 1: "Esto es un test", 2: "Hola! Soy lele :)"}
	fmt.Println(m)

	v := m[1]
	fmt.Println(v)

	// delete key
	delete(m, 1)
	// esto me va a devolver algo "raro" porque ya borre la key 1, me devuelve un string vacio (que es el valor cero.. ni idea.. se le dice asi)
	v = m[1]
	fmt.Println(v)

	// puedo usar found para chequear si existe
	_, found := m[1]
	fmt.Println(found) // devuelve false porque no lo encuentra

	// puedo updatear un value de un map ingresando por su key y asignando un nuevo valor
	// interesante, si la key no existe te la va a crear... o sea que hace un "upsert"
	m[1] = "Ahora tengo un value!"
	value, found := m[1]
	if found {
		fmt.Println(value)
	}
}

func testTypes() {
	// Ingrediente es una estructura que representa los ingredientes de un producto
	type Ingrediente struct {
		nombre          string
		porcentajeGrasa float32
	}

	// Producto es la representacion de un producto comestible
	type Producto struct {
		precio       float32
		nombre       string
		descripcion  string
		marca        string
		ingredientes []Ingrediente
	}

	ing1 := Ingrediente{nombre: "Leche fermentada", porcentajeGrasa: 35.2}
	ing2 := Ingrediente{nombre: "Cereal de trigo", porcentajeGrasa: 4.2}
	yogourt := Producto{precio: 70, nombre: "Yogurisimo con cereales", descripcion: "Yogourt a base de leche fermentada. Trae como acompañamiento cereales", marca: "La Serenisima", ingredientes: []Ingrediente{ing1, ing2}}

	// %v => verbose
	// %#v => muy verbose jajaja
	fmt.Printf("%#v", yogourt)
}
