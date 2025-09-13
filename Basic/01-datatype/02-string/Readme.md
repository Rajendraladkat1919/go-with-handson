| Feature | `string`                      | `rune`                    | `byte`                    |
| ------- | ----------------------------- | ------------------------- | ------------------------- |
| Meaning | Text (sequence of characters) | A Unicode character       | An ASCII/UTF-8 byte       |
| Quotes  | Double (`"A"`)                | Single (`'A'`)            | Single (`'A'`)            |
| Size    | Variable (UTF-8 encoded)      | 4 bytes                   | 1 byte                    |
| Type    | `string`                      | `rune` (alias of `int32`) | `byte` (alias of `uint8`) |
