# Basic
| Type                           | Description                           | Example                    |
| ------------------------------ | ------------------------------------- | -------------------------- |
| `int`                          | Integer (platform-dependent size)     | `var x int = 10`           |
| `int8`/`int16`/`int32`/`int64` | Signed integers                       | `var a int8 = -128`        |
| `uint`                         | Unsigned integer                      | `var u uint = 20`          |
| `float32`, `float64`           | Floating-point numbers                | `var f float64 = 3.14`     |
| `complex64`, `complex128`      | Complex numbers                       | `var c complex64 = 1 + 2i` |
| `bool`                         | Boolean (`true` or `false`)           | `var b bool = true`        |
| `string`                       | Text (UTF-8 encoded)                  | `var s string = "Go"`      |
| `byte`                         | Alias for `uint8` (for raw bytes)     | `var b byte = 'A'`         |
| `rune`                         | Alias for `int32` (Unicode character) | `var r rune = '🦊'`        |


#  complex
| Type      | Description                     | Example                              |
| --------- | ------------------------------- | ------------------------------------ |
| `array`   | Fixed-size sequence of elements | `var a [3]int = [3]int{1, 2, 3}`     |
| `slice`   | Dynamic-size list               | `s := []int{1, 2, 3}`                |
| `map`     | Key-value store                 | `m := map[string]int{"a": 1}`        |
| `struct`  | Custom type with named fields   | `type Person struct { Name string }` |
| `pointer` | Reference to a memory address   | `var p *int = &x`                    |
