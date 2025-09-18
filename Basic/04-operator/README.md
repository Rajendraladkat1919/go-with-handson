kinds of operator in go language.

1. comparision
2. arthemetic
3. Assignment
4. bitwise
5. logical

- Comparison operators only work with comparable types.
- You cannot compare a slice, map, or function using == or != (except comparing them to nil).
- Structs can be compared if all their fields are comparable.

Types that are comparable:

You can use comparison operators with:

Basic types: int, float64, string, bool

Pointers

Channels

Arrays (if the element types are also comparable)

Structs (if all fields are comparable)