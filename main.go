package main // Let's go compiler knows that this program is supposed to be run as a standalone program, instead of a library which will be imported in other programs

import (
	"fmt" // Import the formatting package from the standard library
)

func getCords() (x, y int) { // Named returns = These named vairables are automatically declared at the top of the function body. These should be used for documenting short functions only
	x, y = 5, 6;
	return; // Naked return = The latest value of named variables declared at the top of function are returned

	// return 9, 10; ==> Explicit return
}

func main() { // =====> ENTRY POINT OF THE PROGRAM
	fmt.Println("Starting the Textio server")

	// Sad way of declaring variables
	var mySkillIssues int
	fmt.Println(mySkillIssues)

	// Goated way using walrus operator (:=).
	// It declares and assigns value in the same line along with automatic type inference based on value
	aiSkillIssues := 1000

	fmt.Println(aiSkillIssues)

	const firstName = "Sahil"
	const lastName = "Jassal"

	x, y := getCords();
	fmt.Printf("X: %d, Y: %d\n", x, y)

	printHelloWorld();
}

func printHelloWorld() {
	defer fmt.Println("Two times") 
	defer fmt.Println("World") // this function's execution is defer to just before enclosing function returns

	fmt.Println("Hello")
}
