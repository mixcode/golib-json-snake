# golib-json

A modified go "encoding/json" package that may writes struct field names in _snake\_case_,  _lowerCamelCase_, and _lowercase_.

## Snake-cased marshaling and unmarshalling.

This package is almost-compatible with Go standard encoding/json package, except for a few additional functions.
Most important functions added are `MarshalSnakeCase()` and `UnmarshalSnakeCase()`. As the name implies, `MarshalSnakeCase()` converts Go's _CamelCase_ named members to _snake\_case_ named JSON field members. `UnmarshalSnakeCase()` is vice versa.


## lowerCamelCase / lowercased marshalling/unmarshalling

function `MarshalLowerCamelCase()` and `UnmarshalLowerCamelCase()` handles JSON field names in lowerCamelCase, that means each word is capitalized except for the first word.
Also, function `MarshalLowerCase()` and `UnmarshalLowerCase()` handles JSON field names in lowercase, that means all words are lowercased and concatenated without spaces.


### Example

```go
package main

import (
	"fmt"

	json "github.com/mixcode/golib-json-snake"
)

func main() {
	type st struct {
		Name          string // JSON object name will be "name"
		AnotherName   string // JSON object name will be "another_name"
		ExplicitCamel string `json:"ExplicitCamel"` // If a name tag is explicitly set, the name will be used as-is
		ExplicitSnake string `json:"explicit_snake"`
		Empty         int
	}

	in := st{"name", "another_name", "camel", "snake", 0}

	// encode with snake case
	encoded, err := json.MarshalSnakeCase(&in, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))

	// Output:
	// {"name":"name","another_name":"another_name","ExplicitCamel":"camel","explicit_snake":"snake","empty":0}
}
```

## Decoding of pre-set structs in "map[string]any" variables

Standard Go encoding/json package can unmarshal any JSON string into `map[string]any` type. In this case, child objects are also unmarshalled to `map[string]any` type.

In this package, child objects could be unmarshalled into proper structs if they are pre-placed to the map.


### Example
```go
package main

import (
	"fmt"

	json "github.com/mixcode/golib-json-snake"
)

func main() {
	// a custom struct
	type myStruct struct {
		Name string
		Yes  bool
	}

	m := make(map[string]any) // A universal map to decode a JSON
	m["payload"] = myStruct{} // If we know a exact type of an object, we can pre-set a receiver for the object

	// the 'payload' member in supplied json will be decoded to myStruct placed at m["payload"]
	err := json.UnmarshalSnakeCase([]byte(`{"ok":true,"payload":{"name":"mixcode","yes":true}}`), &m)
	if err != nil {
		panic(err)
	}

	// print the result
	var p myStruct = m["payload"].(myStruct)
	fmt.Printf("%v", p)
	// Output:
	//{mixcode true}
}
```



