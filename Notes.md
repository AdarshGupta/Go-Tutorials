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

## Loops: FOR

Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:
 - the init statement: executed before the first iteration
 - the condition expression: evaluated before every iteration
 - the post statement: executed at the end of every iteration
```Go
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
```
- The init and post statements are optional.
- **For is Go's "while"**: you can drop the semicolons: C's while is spelled for in Go.
```Go
	sum := 1
	for sum < 1000 {
		sum += sum
	}
```
 - **Forever:** If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
 Ex: `for { .. }`

## IF
Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

```Go
if x < 0 {
	return sqrt(-x) + "i"
}
```
- **If with a short statement:** Like for, the if statement can start with a short statement to execute before the condition.
```Go
if v := math.Pow(x, n); v < lim {
	return v
}
```

- **If and else:** Variables declared inside an if short statement are also available inside any of the else blocks.
```Go
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
```

## Switch

 - It runs the first case whose value is equal to the condition expression.
 - Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow.
 - Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
 ```Go
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
 ```
 - Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
 - **Switch with no condition:** Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
 ```Go
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
 ```

## Defer

- A defer statement defers the execution of a function until the surrounding function returns.
- The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
- Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

## Pointers

Go has pointers. A pointer holds the memory address of a value.

- The type *T is a pointer to a T value. 
- Its zero value is nil.
- The & operator generates a pointer to its operand.
- The * operator denotes the pointer's underlying value. This is known as "dereferencing" or "indirecting".
- Unlike C, Go has no pointer arithmetic.

```Go
var p *int
i := 42
p = &i
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

## Structs

- A struct is a collection of fields.

```Go
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// v := Vertex{1, 2}
}

//Output
// >> {1 2}
```

- Struct fields are accessed using a dot.
- Struct fields can be accessed through a struct pointer. To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```Go
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v) // {1000000000 2}
```

```Go
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
}
```

## Arrays

- The type [n]T is an array of n values of type T. 
- The expression `var a [10]int` declares a variable a as an array of ten integers.
- An array's length is part of its type, so arrays cannot be resized.

## Slices

- A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. 
- In practice, slices are much more common than arrays.
- The type `[]T` is a slice with elements of type T. 
- A slice is formed by specifying two indices, a low and high bound, separated by a colon: `a[low : high]`. This selects a half-open range which includes the first element, but excludes the last one.
```Go
var s []int = primes[1:4]
```

- Slices are like references to arrays. A slice does not store any data, it just describes a section of an underlying array.
- Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.

### Slice Literal

A slice literal is like an array literal without the length.

- array literal:
`[3]bool{true, true, false}`
- And this creates the same array as above, then builds a slice that references it:
`[]bool{true, true, false}`
```Go
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
```

### Slice Defaults

- When slicing, you may omit the high or low bounds to use their defaults instead.
- The default is zero for the low bound and the length of the slice for the high bound.

### Slice length and capacity
- The length of a slice is the number of elements it contains.
- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

```Go
s := []int{2, 3, 5, 7, 11, 13}
s = s[:0] // len=0 cap=6 []
s = s[:4] // len=4 cap=6
s = s[2:] // len=2 cap=4
```

### Nil slices
- The zero value of a slice is nil.
- A nil slice has a length and capacity of 0 and has no underlying array. Ex: `var s []int`

## Creating a slice with make

- Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
- The make function allocates a zeroed array and returns a slice that refers to that array: `a := make([]int, 5)  // len(a)=5`
- To specify a capacity, pass a third argument to make:
```Go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

## Slices of slices

Slices can contain any type, including other slices.
```Go
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
```

## Appending to a slice

`func append(s []T, vs ...T) []T`

- The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
- The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
- If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

## Range

The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```Go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
	fmt.Printf("2**%d = %d\n", i, v)
}

// OR

for i, _ := range pow
for _, value := range pow
for i := range pow
```

## Maps

- A map maps keys to values.
- The zero value of a map is nil. A nil map has no keys, nor can keys be added.
- The make function returns a map of the given type, initialized and ready for use.

```Go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

m = make(map[string]Vertex)
m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}
```

## Map literals

- Map literals are like struct literals, but the keys are required.
```Go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```
- If the top-level type is just a type name, you can omit it from the elements of the literal.
```Go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

## Mutating Maps

- Insert or update an element in map m: `m[key] = elem`
- Retrieve an element: `elem = m[key]`
- Delete an element: `delete(m, key)`
- Test that a key is present with a two-value assignment: `elem, ok = m[key]`
- If `key` is in `m`, `ok` is true. If not, `ok` is false. If `key` is not in the map, then `elem` is the zero value for the map's element type.
- Note: If `elem` or `ok` have not yet been declared you could use a short declaration form: `elem, ok := m[key]`

## Function values

- Functions are values too. They can be passed around just like other values.
- Function values may be used as function arguments and return values.

```Go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

hypot := func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

fmt.Println(hypot(5, 12))

fmt.Println(compute(hypot))
```

## Function closures

- Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

- For example, the adder function returns a closure. Each closure is bound to its own sum variable.
```Go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// Output
// 0 0
// 1 -2
// 3 -6
// 6 -12
// 10 -20
// 15 -30
// 21 -42
// 28 -56
// 36 -72
// 45 -90
```

## Methods

- Go does not have classes. However, you can define methods on types.
- A method is a function with a special receiver argument.
- The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
```Go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```
- Remember: a method is just a function with a receiver argument.

- You can declare a method on non-struct types, too.
- You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
```Go
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

## Pointer receivers

- You can declare methods with pointer receivers. This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
- Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
```Go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
}
```

- methods with **pointer receivers** take either a value or a pointer as the receiver when they are called
- Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the Scale method has a pointer receiver.
```Go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
- methods with **value receivers** take either a value or a pointer as the receiver when they are called:

## Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.
- The first is so that the method can modify the value that its receiver points to.
- The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

**Note:** In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

## Interface

- An interface type is defined as a set of method signatures.
- A value of interface type can hold any value that implements those methods.

```Go
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs()) // Error
}
```
**Note:** Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

- A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
- Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
```Go
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```
### Interface values

- Under the hood, interface values can be thought of as a tuple of a value and a concrete type: `(value, type)`
- An interface value holds a value of a specific underlying concrete type. Calling a method on an interface value executes the method of the same name on its underlying type.

### Interface values with nil underlying values

- If the concrete value inside the interface itself is nil, the method will be called with a nil receiver. In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver.

**Note:** Note that an interface value that holds a nil concrete value is itself non-nil.
```Go
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M() // output: <nil>

	i = &T{"hello"}
	describe(i)
	i.M()
}
```

### Nil interface values
- A nil interface value holds neither value nor concrete type. Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

### The empty interface

- The interface type that specifies zero methods is known as the empty interface: `interface{}`
- An empty interface may hold values of any type. (Every type implements at least zero methods.)
- Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

```Go
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## Type assertions

- A type assertion provides access to an interface value's underlying concrete value.
- `t := i.(T)` - This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t If i does not hold a T, the statement will trigger a panic.

- To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
```Go
t, ok := i.(T)
```
- If i holds a T, then t will be the underlying value and ok will be true.
- If not, ok will be false and t will be the zero value of type T, and no panic occurs.

## Type switches

- A type switch is a construct that permits several type assertions in series.
```Go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
**Note:** The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.

## Stringers

- One of the most ubiquitous interfaces is Stringer defined by the fmt package.
```Go
type Stringer interface {
    String() string
}
```
- A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
```Go
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```
## Errors

- Go programs express error state with error values. The error type is a built-in interface similar to fmt.Stringer:
```Go
type error interface {
    Error() string
}
```
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

- Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
```Go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
A nil error denotes success; a non-nil error denotes failure.

```Go
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

## Readers
- The io package specifies the io.Reader interface, which represents the read end of a stream of data. The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

- The io.Reader interface has a Read method:
`func (T) Read(b []byte) (n int, err error)`. Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

The example code creates a strings.Reader and consumes its output 8 bytes at a time.
```Go
import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```

## Images

- Package image defines the Image interface:
```Go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
**Note:** the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package
```Go
import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
```