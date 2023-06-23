/* Go supports embedding of structs and interfaces to express a more seamless composition of types */

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base.
// An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base’s fields directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	// Since container embeds base,
	// the methods of base also become methods of a container.
	// Here we invoke a method that was embedded from base directly on co.
	fmt.Println("describe:", co.describe())

	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	type describer interface {
		describe() string
	}

	// Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
