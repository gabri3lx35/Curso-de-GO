package main
import (
	"fmt"
)

var y = "Opa, bom dia"

func aula2(){
		x := 10
		fmt.Printf("x: %v, %T\n", x,x)
		fmt.Printf("y: %v, %T\n", y,y)

		x,z := 20,30
		fmt.Printf("x: %v, %T\n", x,x)
		fmt.Printf("z: %v, %T\n", z,z)
}


