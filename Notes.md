# Go Notes

## Packages

 - Programs start running in package main.
 - By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

## Imports

But it is good style to use the factored import statement.
Eg: 
```Go
import (
	"fmt"
	"math"
)
```
This code groups the imports into a parenthesized, "factored" import statement.

## Exported names

In Go, a name is exported if it begins with a capital letter.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

## Functions

 - A function can take zero or more arguments.
 - Notice that the type comes after the variable name.
 - When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
 - A function can return any number of results. Ex: 
 ```Go
 func swap(x, y string) (string, string) { ... }
 ```
 - **Named return values**: Go's return values may be named. If so, they are treated as variables defined at the top of the function. A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

 ## Variables

 - The var statement declares a list of variables; as in function argument lists, the type is last.
 - A var statement can be at package or function level.

 Ex: `var c, python, java bool`

 - Variables with initializers: A var declaration can include initializers, one per variable. If an initializer is present, the type can be omitted.

 Ex: `var i, j int = 1, 2`

 - Short variable declarations: Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

 Ex: `c, python, java := true, false, "no!"`

 ## Basic Types

Go's basic types are
```Go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

- Also that variable declarations may be "factored" into blocks, as with import statements.
```Go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```
```Go
fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) 
//Output -  Type: bool Value: false
```
**Note:** The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.

**Zero values:** Variables declared without an explicit initial value are given their zero value.

The zero value is:
 - 0 for numeric types,
 - false for the boolean type, and
 - "" (the empty string) for strings.

 **Type conversions:** The expression T(v) converts the value v to the type T. Unlike in C, in Go assignment between items of different type requires an explicit conversion.
 ```Go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// OR

i := 42
f := float64(i)
u := uint(f)
 ```

**Constants:** Constants are declared like variables, but with the const keyword. Constants can be character, string, boolean, or numeric values. Constants cannot be declared using the := syntax.

Ex: `const Pi = 3.14`
