/*

When its not working:

case 1: 

You can’t use := if all variables on the left are already declared in the same scope.

The short declaration must introduce at least one new variable.
```
var x int = 5
x := 10 // ❌ ERROR: no new variables on left side of :=

```

case 2: correct way:
var x int = 5
x = 10 // ✅ OK: reassigning existing variable

case 3: Mixed declaration
You can mix new and existing variables in a short declaration.	
```
x := 5
x, y := 10, 20 // ✅ OK: `y` is new, so short declaration is allowed
```
case 4: Different scopes
Short variable declarations are scoped to the nearest enclosing function, block, or package level.
```
x := 5
if true {
	x := 10 // ✅ OK: new `x` in inner scope
	fmt.Println(x) // Outputs: 10
}
fmt.Println(x) // Outputs: 5

x := 5 // ❌ ERROR: non-declaration statement outside function body

var x = 5 // OK at package level

case 5: Not for package-level declarations

- When declaring methods or struct fields (you must use var, const, or explicit types)
```
type Person struct {
    name string
    age  int
}

// ❌ Not allowed: you can't use := to declare struct fields.
```
- With constants (you must use const, not :=)

```
const pi = 3.14 // ✅
pi := 3.14      // ❌ ERROR: cannot use short declaration for constants

```

Recap: Long story short

- Inside a function

- At least one new variable on the left

- Not used for constants

- Not at package level