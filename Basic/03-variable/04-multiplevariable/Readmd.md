| Feature                  | `:=`                           | `=`                              |
| ------------------------ | ------------------------------ | -------------------------------- |
| Declares a new variable  | ✅ Yes                          | ❌ No                             |
| Reassigns existing var   | ✅ (if at least one var is new) | ✅ Yes                            |
| Scope                    | Function scope only            | Anywhere                         |
| Allowed at package level | ❌ No                           | ✅ Yes                            |
| Requires type            | No (type is inferred)          | Not always (can infer or define) |
