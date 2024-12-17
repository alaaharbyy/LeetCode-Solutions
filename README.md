# LeetCode Solutions Repository

This repository contains solutions to LeetCode problems, organized by topic. Solutions are implemented in a modular way to make it easy to add new questions and topics in the future.

## Structure
Currently, the repository contains solutions for the following categories:

1. **Strings**
2. **Arrays**

Each topic has its own file, where solutions to relevant questions are implemented.

## How to Use
To control which question gets answered, modify the **`question` variable** in the `main` function. Assign it the appropriate value for the question you want to run.

### Example:
Below is the structure of the `main` function:

```go
package main

import (
	"fmt"
)

func main() {
	question := 1 // Replace with the target question number

	// Strings questions
	s1 := "dd" // Replace s1 and s2 to test other cases
	s2 := "ahbdge" 
	fmt.Println(StringSolutions(s1, s2, question))

	// Arrays questions
	nums := []int{1, 2, 0, 3, 4, 0, 0, 5}  // Replace nums and n to test other cases
	n := 2
	fmt.Println(ArraySolutions(nums, n, question))
}
```

### Steps:
1. Set the `question` variable to the desired question number.
2. Input the required parameters for the specific question under strings or arrays.
3. Run the program to view the solution.

Stay tuned for updates!

---
**Happy Coding!** 🚀
