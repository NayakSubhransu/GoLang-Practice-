package main

import "fmt"

// func main() {
//     name := "Subhransu"
//     age := 25
//     // fmt.Println("Name:", name, "Age:", age)
// 	// fmt.Print("Hello ")
//     // fmt.Print("World")

//    fmt.Printf("Name: %s \nAge: %d\n", name, age)
// }

// %d	integer
// %s	string
// %f	float
// %t	boolean

func main() {
    var age int = 25
    var name string = "Subhransu"
    var pi float64 = 3.14

	var a, b, c int = 1, 2, 3
    x, y := 10, "Hello"

    fmt.Println(age, name, pi)
    fmt.Println(a, b, c)
    fmt.Println(x, y)

}