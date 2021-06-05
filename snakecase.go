package json

// convert CamelCase name to camel_case
func toSnakeCase(key string) string {
	return string(toSnakeCaseByte([]byte(key)))
}

// convert snake_case name to SnakeCase
func toCamelCase(key string) string {
	return string(toCamelCaseByte([]byte(key)))
}

func toSnakeCaseByte(key []byte) []byte {
	buf := make([]byte, len(key)*2)
	j := 0
	for i := 0; i < len(key); i++ {
		c := key[i]
		if c >= 'A' && c <= 'Z' {
			if i != 0 {
				buf[j] = '_'
				j++
			}
			c = c + 0x20 // convert to lowercase
		}
		buf[j] = c
		j++
	}
	return buf[:j]
}

func toCamelCaseByte(key []byte) []byte {
	buf := make([]byte, len(key))
	j := 0
	toUpper := true
	for i := 0; i < len(key); i++ {
		c := key[i]
		if c == '_' {
			toUpper = true
			continue
		}
		if toUpper && c >= 'a' && c <= 'z' {
			c = c - 0x20
		}
		toUpper = false
		buf[j] = c
		j++
	}
	return buf[:j]
}
