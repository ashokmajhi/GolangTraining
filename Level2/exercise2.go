/*
Using the following operators, write expressions and assign their values to variables:
g.	==
h.	<=
i.	>=
j.	!=
k.	<
l.	>
Now print each of the variables.

*/
package main

import "fmt"

func main() {
	a := (13 == 13)
	b := (13 <= 16)
	c := (13 >= 16)
	d := (14 != 16)
	e := (11 < 23)
	f := (55 > 23)
	fmt.Println(a, b, c, d, e, f)
}
