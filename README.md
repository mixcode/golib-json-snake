# golib-json

A modified go "encoding/json" package writes struct field names in snake_case.

## Snake-cased marshaling and unmarshalling.

This package is almost-compatible with Go standard encoding/json package, except for two additional functions.
The added functions are "MarshalSnakeCase()" and "UnmarshalSnakeCase()".
As the name implies, MarshalSnakeCase() converts Go's CamelCase named members to snake_case named JSON object members.
And UnmarshalSnakeCase() decodes JSON snake_cased objects into Go CamelCase members.

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
		ExplicitCamel string `json:"ExplicitCamel"` // If tag is explicitly set, the name will be used
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

## Pre-set struct to "map[string]any" variable.
Standard Go encoding/json package can unmarshal any JSON string into "map[string]any" type. In this case, child objects are also unmarshalled to "map[string]any" type.

But in this package, child objects could be unmarshalled into proper structs if they are pre-placed to the map.


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

	m := make(map[string]any) // A map to unmarshal JSON into
	m["payload"] = myStruct{} // place a receiver variable to the key "payload"

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



