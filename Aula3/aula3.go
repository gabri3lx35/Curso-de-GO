package main

import "fmt"

//tipos em Go são estaticos
//tipos de dados Primitivos (int, string, bool)
//valor 0
var a int
var b float64
var c string
var d bool
func aula3(){
  fmt.Printf("%v, %T\n", a,a )
  fmt.Printf("%v, %T\n",b,b )
  fmt.Printf("%v, %T\n",c,c )
  fmt.Printf("%v, %T\n",d,d )	
}

