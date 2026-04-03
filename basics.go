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

// func main() {
//     var age int = 25
//     var name string = "Subhransu"
//     var pi float64 = 3.14

// 	var a, b, c int = 1, 2, 3
//     x, y := 10, "Hello"

//     fmt.Println(age, name, pi)
//     fmt.Println(a, b, c)
//     fmt.Println(x, y)

// }


func main() {
    score := 85

    if score >= 90 {
        fmt.Println("A grade")
    } else if score >= 80 {
        fmt.Println("B grade")
    } else {
        fmt.Println("Below B")
    }

    // Go's unique: if with initialization statement
    // The variable `result` only exists inside this if block
    if result := score * 2; result > 150 {
        fmt.Println("High score:", result)
    } else {
        fmt.Println("Normal score:", result)
    }
    // result is NOT accessible here — great for scoping
}

