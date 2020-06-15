// names

package primitives

type NameBundle struct {
	FirstName string
	LastName  string
}

func Name() NameBundle {

	test := NameBundle{"John", "Smith"}
	return test
}
