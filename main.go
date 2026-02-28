package main // Let's go compiler knows that this program is supposed to be run as a standalone program, instead of a library which will be imported in other programs

import "fmt" // Import the formatting package from the standard library

func main() { // =====> ENTRY POINT OF THE PROGRAM
	fmt.Println("Starting the Textio server")

	// Sad way of declaring variables
	var mySkillIssues int
	fmt.Println(mySkillIssues)

	// Goated way using walrus operator (:=). 
	// It declares and assigns value in the same line along with automatic type inference based on value
	aiSkillIssues := 1000

	fmt.Println(aiSkillIssues)
}
