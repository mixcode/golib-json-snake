package json

import (
	"fmt"
)

func ExampleMarshalSnakeCase() {
	type st struct {
		Name          string
		AnotherName   string
		ExplicitCamel string `json:"ExplicitCamel"`
		ExplicitSnake string `json:"explicit_snake"`
		Empty         int
	}

	in := st{"name", "another_name", "camel", "snake", 0}

	// encode with snake case
	encoded, err := MarshalSnakeCase(&in, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))

	// encode with snake case + omit empty fields
	encoded, err = MarshalSnakeCase(&in, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))

	// Output:
	// {"name":"name","another_name":"another_name","ExplicitCamel":"camel","explicit_snake":"snake","empty":0}
	// {"name":"name","another_name":"another_name","ExplicitCamel":"camel","explicit_snake":"snake"}
}
