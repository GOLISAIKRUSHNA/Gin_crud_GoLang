package main

import (
	"fmt"
	"os"
)

func main() {

	const name string = "goli"
	const fullname = "saikrushna goli"

	var age = 23
	fmt.Println(name)
	fmt.Println(age, fullname)

	fmt.Println(os.Environ())
	fmt.Println(os.Getenv("Shell"))
	os.Setenv("DUMMY", "dummy")
	fmt.Println(os.Getenv("DUMMY"))

}
