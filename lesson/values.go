package lesson

import "fmt"

/* Go has various value types including strings, integers, floats, booleans, etc... */

/* `Values` function: Print the values of the variables like strings, integers, floats, booleans, etc... */
func Values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/*`Values` function output:
- golang
- 1+1 = 2
- 7.0/3.0 = 2.3333333333333335
- false
- true
- false
*/
