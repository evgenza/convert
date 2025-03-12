# Convert — A Handy Go Library for Data Conversion

Convert is a lightweight Go library that provides convenient conversion functions between various types, formats, and encodings. It wraps the standard `strconv`, `encoding/base64`, `encoding/hex`, `encoding/json`, and `encoding/xml` packages to make conversions more intuitive.

## Features

- Convert strings to/from integers, floats, booleans.
- Encode/decode base64 and hex.
- Serialize/deserialize JSON and XML.
- Simple API with clear function names.
- Well-tested with Go's built-in testing framework.

## Installation

```bash
go get github.com/evgenza/convert
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/evgenza/convert"
)

func main() {
    i, err := convert.ToInt("123")
    if err != nil {
        panic(err)
    }
    fmt.Println("Int value:", i)

    s := convert.FromInt(99)
    fmt.Println("String value:", s)

    encoded := convert.ToBase64([]byte("hello world"))
    fmt.Println("Base64 encoded:", encoded)
    decoded, _ := convert.FromBase64(encoded)
    fmt.Println("Base64 decoded:", string(decoded))

    jsonStr, _ := convert.ToJSON(map[string]int{"foo": 1, "bar": 2})
    fmt.Println("JSON string:", jsonStr)

    var m map[string]int
    _ = convert.FromJSON(jsonStr, &m)
    fmt.Println("Deserialized map:", m)
}
```

## License

This project is released under the MIT License.
