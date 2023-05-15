// Provides utility functions for the console.

package util

import "fmt"

// ClearConsole clears the console (stdout).
// This should work across all terminals.
func ClearConsole() {
	// This works by an ANSI sequence:
	//
	// \033[H    Move the cursor to the top-left corner of the console
	//
	// \033[2J   Clear the character at the cursor and every character
	//           after it
	fmt.Print("\033[H\033[2J")
}
