package json

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {

	snakes := []string{
		"snake",
		"snake_case",
		"snake_cased_1text_len&gth",
	}
	snakeOk := []string{
		"Snake",
		"SnakeCase",
		"SnakeCased1textLen&gth",
	}
	for i, s := range snakes {
		c := toCamelCase(s)
		if c != snakeOk[i] {
			t.Errorf("CamelCase conversion failed: %s -> %s", s, c)
		}
	}

	camels := []string{
		"Camel",
		"CamelCase",
		"Camel1CasedSt%ring",
	}
	camelOk := []string{
		"camel",
		"camel_case",
		"camel1_cased_st%ring",
	}
	for i, s := range camels {
		c := toSnakeCase(s)
		if c != camelOk[i] {
			t.Errorf("snake_case conversion failed: %s -> %s", s, c)
		}
	}

}

func TestLowerCamelCaseT(t *testing.T) {

	type st struct {
		TestKey struct {
			TestValue int
		}
	}

	j := []byte(`{"testKey":{"testValue":1, "testV2": 2}}`)

	//k := make(map[string]interface{})
	var k st

	err := UnmarshalLowerCamelCase(j, &k)
	if err != nil {
		t.Fatal(err)
	}
	if k.TestKey.TestValue != 1 {
		t.Fatalf("incorrect decoded value")
	}

	m, err := MarshalLowerCamelCase(k, false)
	if err != nil {
		t.Fatal(err)
	}

	if string(m) != `{"testKey":{"testValue":1}}` {
		t.Fatalf("incorrect encoding value")
	}
}
