package json

import (
	ej "encoding/json"
	"testing"
)

func TestSnakeEncode(t *testing.T) {
	var err error

	type st struct {
		Name          string
		AnotherName   string
		ExplicitCamel string `json:"ExplicitCamel"`
		ExplicitSnake string `json:"explicit_snake"`
		Empty         int
	}

	in := st{"name", "another_name", "camel", "snake", 0}
	ok1 := `{"name":"name","another_name":"another_name","ExplicitCamel":"camel","explicit_snake":"snake","empty":0}`
	ok2 := `{"name":"name","another_name":"another_name","ExplicitCamel":"camel","explicit_snake":"snake"}`
	okJ := `{"Name":"name","AnotherName":"another_name","ExplicitCamel":"camel","explicit_snake":"snake","Empty":0}`

	// encode with snake case
	encoded, err := MarshalSnakeCase(&in, false)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != ok1 {
		t.Fatal("encoded data not match")
	}
	var out1 st = st{"a", "b", "c", "s", 1}
	err = UnmarshalSnakeCase(encoded, &out1)
	if err != nil {
		t.Fatal(err)
	}
	if in != out1 {
		t.Fatal("decoded value not match")
	}

	// snake case + ignore all empty
	encodedI, err := MarshalSnakeCase(&in, true)
	if err != nil {
		t.Fatal(err)
	}
	if string(encodedI) != ok2 {
		t.Fatal("encoded-ignore data not match")
	}
	var out2 st = st{"a", "b", "c", "s", 0}
	err = UnmarshalSnakeCase(encodedI, &out2)
	if err != nil {
		t.Fatal(err)
	}
	if in != out2 {
		t.Fatal("decoded value not match")
	}

	// encoding/json
	encodedJ, err := ej.Marshal(&in)
	if err != nil {
		t.Fatal(err)
	}
	if string(encodedJ) != okJ {
		t.Fatal("encoded-ignore data not match")
	}
	var outJ st = st{"a", "b", "c", "s", 0}
	err = UnmarshalSnakeCase(encodedJ, &outJ)
	if err != nil {
		t.Fatal(err)
	}
	if in != outJ {
		t.Fatal("decoded value not match")
	}

}
