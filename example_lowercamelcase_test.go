package json

import (
	"fmt"
)

func ExampleMarshalLowerCamelCase() {
	type st struct {
		Name          string
		AnotherName   string
		OneMoreName   string
		ExplicitCamel string `json:"ExplicitCamel"`
		ExplicitSnake string `json:"explicit_snake"`
		Empty         int
	}

	s1 := st{"name", "anotherName", "", "Camel", "snake", 0}

	// encode with snake case
	encoded, err := MarshalAs(&s1, LowerCamelCase, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))

	// decode json
	var s2 st
	err = UnmarshalLowerCamelCase(encoded, &s2)
	if err != nil {
		panic(err)
	}

	// encode with snake case + omit empty fields
	encoded, err = MarshalAs(&s2, LowerCamelCase, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))

	// Output:
	// {"name":"name","anotherName":"anotherName","oneMoreName":"","ExplicitCamel":"Camel","explicit_snake":"snake","empty":0}
	// {"name":"name","anotherName":"anotherName","ExplicitCamel":"Camel","explicit_snake":"snake"}
}
