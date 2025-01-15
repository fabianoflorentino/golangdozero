// Package fundamentals_of_language provides examples and demonstrations of fundamental concepts in the Go programming language.
// This package includes examples of creating and initializing fixed-length arrays, as well as other basic language features.
package fundamentals_of_language

import (
	"fmt"

	"github.com/fabianoflorentino/golangfromzero/pkg/trim"
)

// exampleArrayWithLength contains a string literal that demonstrates the creation and initialization
// of fixed-length arrays in Go. It includes three examples:
// Example 1: Declares an array of 5 strings, assigns a value to the first element, and prints it.
// Example 2: Declares and initializes an array of 5 strings with specific values.
// Example 3: Declares and initializes an array of integers using the ellipsis syntax to infer the length.
var (
	exampleArrayWithLength = `
// arrayWithLength demonstrates the creation and initialization of a fixed-length array in Go.
// It declares an array of 5 strings, assigns a value to the first element, and prints it.

Example 1:

var arr1 [5]string
arr1[0] = "a"

Example 2:

arr2 := [5]string{"a", "b", "c", "d", "e"}

Example 3:

arr3 := [...]int{1, 2, 3, 4, 5}
`
)

// ArraysAndSlices demonstrates the usage of fixed-length arrays in Go.
// It initializes a new trim object, prints a message indicating the use of fixed-length arrays,
// and then prints the trimmed string representation of an example array with a fixed length.
// Finally, it calls the arrayWithLength function to further illustrate array operations.
func ArraysAndSlices() {
	trim := trim.New()

	println("\nWith fixed-length arrays:\n")
	println(trim.String(exampleArrayWithLength))
	arrayWithLength()
}

// arrayWithLength demonstrates the creation and initialization of a fixed-length array in Go.
// It declares an array of 5 strings, assigns a value to the first element, and prints it.
func arrayWithLength() {
	println("\nResults:")

	var arr1 [5]string
	arr1[0] = "a"
	println("\nResult Example 1: ", arr1[0])

	arr2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("\nResult Example 2: ", arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("\nResult Example 3: ", arr3)
}
