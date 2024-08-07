package main

import "fmt"



func main() {

	//pass by value
	// age:=20
	// fmt.Println("Before calling func",age)
	// name(age)
	// fmt.Println("after calling func",age)

	// var argu string="this is string"
	// fmt.Println(argu)
	// val:=&argu
	// fmt.Println(val)//referencing

	// fmt.Println(*val)//dereferencing


	var argu string="content"
	fmt.Println("Before the calling ",argu)


	//pass by reference
	name(&argu)

	fmt.Println("after the calling ",argu)


	







}

func name( argu *string){
	*argu="main content"
	
	// fmt.Println(argu)
	




}


