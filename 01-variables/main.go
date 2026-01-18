package main
import "fmt"
func main() {
	var a int
	var b string
	fmt.Println("Default value of int a:", a)
	fmt.Println("Default value of string b:", b)

	//initialize variables
	a = 10 // assign value to a
	fmt.Println(a)


	//CREATE AND INITIALIZE variable
    var c = 10
	var d = "cokiecode"

	fmt.Println(c)
	fmt.Println(d)


	// Declaring multiple variables
	var e, f int
	fmt.Println(e, f)
	e = 10
	f = 100
	fmt.Println(e, f)
	// Creating and initializing multiple variables
	var g, h = 10, 100
	fmt.Println(g, h)

	// Variable Declaration Block
	var (
		i = 10
		j = "mohammad"
		k = 1000
	)

	fmt.Println(i, j, k)

	//Short variable declaration
	var l int
	var m int
	l, m, n := 10, 20, 30
	fmt.Println(l, m, n)




}