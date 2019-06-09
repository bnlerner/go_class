/*
EXPLANATION OF RACE CONDITION
A race condition is a situation where the output of a program is dependent on the order of operations in a program.
This output becomes an error when the result is not expected by the programmer. Race conditions usually occur when
multiple threads are accessing the same shared resource. When the order of operations on this resource is controlled
by a thread scheduling algorithm and it can run these threads in any order concurrently or by a multi threaded system
it can lead to unexpected results.

EXPLANATION OF THIS SPECIFIC RACE CONDITION
In the submitted code, the Race condition is due to the increment operation and
print operation using a shared variable num. These operations run concurrently but in whatever order the system thread scheduler
decides. This means the value printed could be 0 or 1 depending on the order of when the increment and print command is executed.

*/

package main

import (
	"fmt"
)

func main() {

	var num int

	go func() { num++ }()

	go func() { fmt.Print("\nnum", num) }()

	fmt.Scanln()
	fmt.Print("done")
}
