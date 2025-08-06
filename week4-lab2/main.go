package main

import (
	"fmt"
)

func main() {

	//var name string = "Supachok"
	var age int = 21

	email := "aintharasopa_s@silpakorn.edu"
	gpa := 3.21

	firstName, lastName := "Supachok", "Aintharasopa"

	fmt.Printf("Name: %s %s,age %d,email %s GPA %.2f\n", firstName, lastName, age, email, gpa)
}
