package main

import (
	"fmt"
	"reflect"
)

func main() {

	var fname, lname, name string = "goli", "saikrushna", "purushottam"
	fmt.Println(fname, lname, name)

	var (
		namee string
		agee int=20
	)
	
	fmt.Println(namee,agee)

	 ab:=12
	 abc:="nana"
	 
	 fmt.Println(ab,abc)
	 fmt.Println(reflect.TypeOf(ab),reflect.TypeOf(abc))

	 integer:=10
	 floater:=20.000000000000000
	 c:=int(floater)+integer
	 fmt.Println(c)




	



}
