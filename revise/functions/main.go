package main
import "fmt"

func name(fname, lname string) (string, string) {
	return fname, lname
}
func sum(arr... int) (int) {
	sum:=0
	for _,j :=range arr{
		sum=sum+j
	}

	return sum
}
func main() {

	fname, lname := name("saikrushna", "goli")
	fmt.Println(fname, lname)

	res:=sum(1,2,3,4,5,7)
	fmt.Println(res)


}

