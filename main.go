package main

import "fmt"

type auto struct {
	nombre string
	modelo string
}

//no funciona
type agencia struct {
	autos := make([]auto, 3)
}

func (a *auto) print() {
	for {
		fmt.Println(autos)
	}
}

func (a agencia) addauto(c auto){
	autos = append(inlineSlice, c)
}

func main() {
	a := &auto{nombre: "BMW", modelo: "135i"}
	a.print()
}
