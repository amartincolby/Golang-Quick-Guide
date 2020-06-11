![Golang Logo](Go_Logo_Blue.svg)

# Golang Quick Guide
This document was inspired by Learn X in Y.

The purpose of this document is primarily a tool for self-edification. By writing and teaching, I better learn. This is in contrast to my ReasonML Quick Guide where I both wanted to learn and also spread the ReasonML gospel. Go already has the gospel (The GOspel, perhaps?) being spread and heard. Still, as with most programming languages, much of the documentation available is poorly written with nary a drop of humanism in how it is delivered. I hope to deliver that humanism. Further, and perhaps I am alone in this feeling, but I hate overly complex tutorials. I prefer one-page guides that allow me to easily search, cross-reference, and flow from point to point without having to jump between pages.

# Further Information

- [Official Go Documentation](https://golang.org/doc/)
- [Golang University Playlists](https://www.youtube.com/channel/UCP67S6tmE0xv62yQLg9hnmg/playlists)
- [Go Language Specification](https://golang.org/ref/spec) (I'm not joking. You can read the entire spec in a day or two.)

# Let's Go
Go is a good language. It is especially satisfying for those in the industry who despise the pandering, bloated direction that many languages have taken in the past two decades. Java, JavaScript, C#, and C++ have all become colossal, with specs so large that no one could possible learn them entirely. Go rejects almost all of the syntactic and semantic progress made in the past thirty years. It is lean, simple, and firmly rooted in the legacy of C, both semantically and syntactically. Its stated goal was to be a better C++. Go is also attractive because its simplicity and rigid style recommendations means that developers cannot be _clever_. Cleverness is the drug of the junior developer and the bain of the senior developer. The best code is the simplest, most easily-read code, and Go actively prevents anything but.

That said, while it is enticing for those who look down upon tool-happy developers who unquestioningly accept the latest developments in a language, this semantic dogmatism also means that the problems of C in 1985 are essentially untouched. This helps to explain oddities such as the language allowing things that are explicitly discouraged. This makes Go easy to understand, but it also means that a Go program can fail. Go is strictly-typed but it is an unsafe language. This sort of unsafety can be annoying for those coming from more academic languages. The speed of Go cannot be impugned, though. It is the fastest garbage-collected language, and despite some truly admirable efforts from the Java world, this crown will likely remain upon Go's head.

Go was designed to build n-tier systems. It is fast enough to be used for most anything, but that is its intended use. To further this, Go contains two semantic constructs meant to facilitate building n-tier applications: Goroutines and channels. Both are language level constructs intended to enable better asynchronicity across a system. Aside from them, though, Go is C, so it should be easily and immediately understandable to most in the programming world.

The below text is valid Go code. It is a copy of the code in the main.go file.

``` golang
/* Comment blocks start with slash-star,
   and end with star-slash. */
// Line comments begin with a double slash.

/* Most IDE's include the standard Go style guide which will auto-format your
code when possible. The style guide is highly-opinionated and will go so far
as to remove unused imports on file save. */

/*----------------------------------------------
 * Build tags
 *----------------------------------------------
 */

/* Go includes variable builds at the language level. By this, I mean that Go
source files can be tagged to only be included in a compilation if certain flags
are present when the compile command is run. These build tags are conveniently
called "build tags" and are a special double-slash comment that is recognized by
the compiler. This enables developers to have builds specifically for develop,
staging, and production environments without having to worry about managing
those arguments programmatically. */

/*   // +build dev, stage   */

/* In the above build tag, this file will only be compiled if the flags for
either dev or stage are included in the build command, such as
`go build -tags="stage"`. Since this obviously does not make sense on the main
file, the star-slash prevents the build tag comment from being read. Further, in
most files, the list of build tags would be at the very top.

Tags can be used to identify anything, not simply dev or prod builds. Entire
experiments can be hidden behind a tag. Or if a project is practicing high-
frequency merging, unfinished features can be kept behind an experimental tag.
In essence, different programs can exist in the same code base. */

/*----------------------------------------------
 * Packages and imports
 *----------------------------------------------
 */

/* All files must begin with a package declaration. This is the identifier that
is used by other files for import and can be shared by many files, all of which
will be included when another file imports that package. For example, if
multiple files all declare `package kwyjibo` at the top of the file, when other
files import `kwyjibo`, they will have access to all files with that package
declaration. Within a file, all names that begin with capital letters will be
exported. Names beginning with a lowercase letter is private and inaccessible
outside of that file.

The only reserved name for a package is `main` because this indicates that the
file is intended for immediate execution as opposed to a package that can be
imported by other files. The `main` package will be the root of any and all Go
projects. Since this one-page guide must be a valid Go file, its package
declaration must be `main`. */

package main

/* The `import` verb declares the external packages used in this file. The below
syntax is called "factored" importing, but it's really just shorthand instead of
using `import` repeatedly. The list of packages is not a tuple. To put imports
on a single line, include them in a semicolon-delimited list. This syntax is
technically correct but will be auto-converted to the standard line-delimited
list by most formatters. */

import (
	"fmt" // A package in the Go standard library.
	// Math library with local alias m.
	// Yes, a web server!
	// OS functions like working with the file system
)

/*----------------------------------------------
 * Variables, pointers, and functions
 *----------------------------------------------
 */

/* Variables in Go are conceptually similar to variables in C, to wit they are
memory allocations. By default, variable declaration and use relies on
references by name, meaning that the runtime can put the value attached to that
name anywhere in memory. Names must begin with a letter or underscore. Names
using underscores are discouraged. Variables can be declared in two ways:
explicit and short-form.

An explicit declaration uses the `var` keyword. Types can be declared or
inferred. If a variable is declared but no value is assigned, the default
value is `nil`. */

var intAllocation int // = nil
var dinnerForOne string = "The same procedure as every year"
var theAnswer = 42

/* explicit declarations can also use factored syntax. */

var (
	partyLikeIts          = 1999
	y2k               int = 2000
	yearWeMakeContact     = 2001
)

/* Short-form declarations work similarly to explicit, un-typed variable
declarations as seen above. They do not use the `var` keyword or a type
declaration. They cannot be used outside of function bodies. */

func setVariables() {
	anInteger := 42
	aString := "Thank you, Thing."

	fmt.Println(anInteger, aString)
}

/* The semantics of short-form declarations are slightly different than regular
declarations. They can re-declare variables that were previously declared in the
same block, whether via explicit declaration, short-form, or function parameters.
Oddly, this re-declaration can only occur when using short-form multi-variable
syntax when at least one of the variables is new. */

func resetVariables(aParameter string) {
	anInteger := 42
	aParameter, anInteger, aNewVariable := "None shall pass!", 2001, true

	fmt.Println(aParameter, anInteger, aNewVariable)
}

/* Variables can be declared outside of function bodies, but already-declared
variables cannot be assigned or reassigned outside of blocks. Below, a function
is used, but any block would work. */

func reserveDinner() {
	dinnerForOne = "Well, I'll do my very best."
}

/*** Pointers ***/

/* While the default allocation returns a named reference, Go supports pointers
to a memory address. In Go, a pointer is a symbol that represents a specific
physical location in memory.  Pointers provide significant performance benefits
but are unsafe. Go does not have pointer arithmetic, meaning that the pointer's
value itself cannot be altered with the usual mathematical operators like `+`
or `-`.

Create a pointer via the `*` prefix on the type declaration. The value contained
at this memory address is `nil` by default. */

var pointerToTheAnswer *int

/* A pointer to an existing value can be created with the address operator `&`.
The address operator is a unary operator that returns the memory address of
a previously declared variable.

Since the setting of a value at an address is an assignment, it cannot happen
outside of a function. The below function exists purely to set the value of
the pointer. */

func setTheAnswer() {
	pointerToTheAnswer = &theAnswer
}

/* When applied to a pointer, the `*` is the dereferencing operator. It is
so-called becaued it extracts the value to which the pointer points, thus
turning it from a reference to a place in memory to an actual value that can
have computation done on it. If a pointer of nil value is dereferenced, a
runtime panic will occur. */

func printTheAnswer() {
	fmt.Println(*pointerToTheAnswer)
}

// Finally, set values into that memory address via the pointer.

func changeTheAnswer() {
	*pointerToTheAnswer = 43
}

/*** Functions ***/

/* Functions obey the widely understood C syntax. All applications must have a
`main()` function that acts as the entry point of a Go project. The main
function must exist within package main. Once the main function fully evaluates,
the Go application will terminate. */

func main() {
	fmt.Println("The Go application has started.")

	fmt.Println("The Go application has finished.")
}

// Function parameters and returns must be explicitly typed.

func whoYouGonnaCall(option int) string {
	var hero string

	if option == 1 {
		hero = "Ghostbusters"
	} else if option == 2 {
		hero = "Ninja Turtles"
	} else if option == 3 {
		hero = "Jem"
	} else {
		hero = "Steve Urkel"
	}

	return hero
}

/* If a function's parameters are the same type, only the last one requires the
type declaration. */

func adder(x, y int) int {
	return (x + y)
}

/* Function return names can be set in the declaration syntax. The function body
can then set the value of the return like any other variable. If this is done,
no return value is specified. */

func loadMovie(option int) (movie string) {
	if option == 1 {
		movie = "Toy Story"
	} else if option == 2 {
		movie = "Dora The Explorer"
	} else if option == 3 {
		movie = "Inside Out"
	} else {
		movie = "Apocalypse Now"
	}
	return
}

/* Functions can return two values. Make note that the function is not returning
a structure like a tuple. It is returning two distinct values that must be
captured at the call site. */

func getFilm(option int) (name string, time int) {
	if option == 1 {
		name = "Titanic"
		time = 2030
	} else if option == 2 {
		name = "The Godfather"
		time = 1900
	} else if option == 3 {
		name = "Black Panther"
		time = 1830
	} else {
		name = "The Room"
		time = 1645
	}
	return
}

var movieName, movieTime = getFilm(2)

/*----------------------------------------------
 * Basic types
 *----------------------------------------------
 */

/*** Type Conversion ***/

/* Go does not support implicit type conversion. Each type has an explicit
conversion function to cast a value of one type as another. */

var (
	anInteger int = 1337
	aFloat        = float64(anInteger)
)

/*** Boolean ***/

var aBoolean bool = true

/*** Integers ***/

/* Go supports manifold integer types, both signed and unsigned. A signed
integer is one where the last bit is used to determine whether the integer is
negative or positive. An unsigned integer is one that uses all available bits to
represent a positive number. As such, unsigned integers can represent double the
number of positive integers as a signed integer. Integers can be broken up with
underscores for formatting and readability purposes.

All numbers shown are largest possible numbers for that type. */

// Unsigned Integers

var (
	uInt1 uint8  = 255
	uInt2 uint16 = 65535
	uInt3 uint32 = 4294967295
	uInt4 uint64 = 1844674407_3709551615
)

// Signed Integers

var (
	sInt1 int8  = 127
	sInt2 int16 = 32767
	sInt3 int32 = 2147483647
	sInt4 int64 = 9223372036854775807
)

// Generic int Type

/* Go also has generic int & uint types whose size is dependent on the platform
for which the application is compiled. As with most languages where this
distinction exists, it is generally best practice to use the generic types
unless explicit knowledge of the integer size is required. When type inference
is used, Go will assume a generic int. */

var (
	uInt uint = 1337
	sInt int  = 8008
)

/*** Octal, Hexidecimal, and Binary ***/

/* Any integer type can be used to represent an octal number, hexidecimal
number, or raw binary number. Octal numbers need simply be prefixed with a zero.
Hex numbers are prefixed with a zero followed by "x". Binary numbers are
prefixed with zero and "b". */

var (
	octalNumber        = 01337
	hexNumber    int32 = 0x1ee7
	binaryNumber       = 0b01101110_01100101_01110010_01100100
)

/*** Floating Point ***/

/* Go supports two types of floats: single and double precision, aka 32-bit and
64-bit. The terms single and double are somewhat out-of-date and should instead
be called binary32 and binary64. They are fully compatible with the IEE-754
standard as found here:

- https://en.wikipedia.org/wiki/Single-precision_floating-point_format
- https://en.wikipedia.org/wiki/Double-precision_floating-point_format

Like integers, floats can be broken up with underscores for formatting and
readability purposes. When type inference is used, Go will assume a float64. */

var (
	float1 float32 = 3.1415927
	float2 float64 = 3.14_1592653589793
	float3         = 3.14
)

/*** Hexidecimal Floats ***/

/* Floats support hexidecimal but requires a declared exponent with the exponent
label "p". */

var hexFloat = 0x3.0p2

/*** Complex Numbers ***/

/* Complex numbers are a combination of two real numbers a and b and the
imaginary number i, where i is a number that satisfies x^2 = -1. A complex
number as expressed is a + bi. Since a complex number is a combination of two
floats, the type is double the combined floats, i.e. float32 = complex64 and
float64 = complex128. When type inference is used, Go will assume complex128.

Complex numbers are useful for computation but will rarely used for n-tier
application development. */

var (
	complex1 complex64  = 13 + 37i
	complex2 complex128 = 20 + 01i
	complex3            = 19 + 99i
)

/* If using previously declared variables, the complex library is used. Note
that the two values used must be of the same float type, in this case float64.*/

var (
	complex4 = complex(float2, float3)
)

/*** Byte & Rune ***/

/* Go does not have a char type to represent single characters. Instead, it
relies on aliases for uint8 and int32, byte and rune. Byte includes the entire
ASCII character set while rune includes all unicode characters. Bytes and runes
are wrapped in single-quotes. When type inhference is used, Go will assume a
rune type. */

var (
	character1   byte = 'a'
	character2   rune = 'b'
	inferredRune      = 'c'
)

/*** String ***/

/* Strings can be declared in two ways: double quotes and backticks. Quotes work
as expected. They cannot span multiple lines and require escaping of certain
characters. The particulars can be found in the Go documentation. */

var aString = "Now for something completely different"

/* Backticks denote raw strings of uninterpreted bytes. They do not require
escape characters and can span multiple lines. The only behavior of note is that
carriage return characters are discarded when the string is interpreted by the
program. */

var aRawString = `Stop that! This is
getting awfully silly!`

/*** Array ***/

/* Arrays are fixed-size sequences of entities of the same type that exist
contiguously in memory. */

var integerArray [4]int

func fillArray() {
	integerArray[0] = 1
	integerArray[1] = 3
	integerArray[2] = 3
	integerArray[3] = 7
}

/* Arrays can be simultaneously declared and populated with curly braces. */

var stringArray = [4]string{"More", "human", "than", "human"}

/* Arrays are one-dimensional, but n-dimensional arrays can be created with
arrays of arrays. */

var cubeArray [4][4][4]string

/*** Slice ***/

/* A slice is the first part of Go that will require some degree of elucidation
for those coming from other lanaguages. Be aware that while languages like
JavaScript and Python have a slice for arrays, it behaves very differently.

A slice is a contiguous section of an underlying array. It does not itself
contain any data. A slice can be seen as a flexible and easy-to-use interface
into an array. Slices provide behavior similar to arrays in more dynamic
languages such as JavaScript and are thus more common than true arrays in
production Go code. Multiple slices can be associated with a single array. A
slice type is specified similarly to an array type, only without a length
between the brackets.

To slice an existing array, provide the starting index inclusive, and ending
index exclusive. Slices can themselves be sliced, with each slice referencing
the same array. */

var (
	anArray       [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	aSlice        []int  = anArray[1:5] // [2, 3, 4, 5]
	sliceOfASlice []int  = aSlice[1:4]  // [3, 4, 5]
)

/* In most cases, free slice literals will be used. Here, no underlying array is
specified before the creation of the slice. Below, the Go automatically runtime
creates an invisible, underlying array for `anotherSlice` to reference. By
default, the slice and the array are of the same size. */

var anotherSlice []int = []int{1, 2, 3, 4, 5}

/* Just as with arrays, multidimensional slices can be created with underlying
arrays automatically created. */

var threeDimensionalSlice [][][]string

/* Slices can also be created with the `make()` command. When used for creating
slices, `make()` accepts three arguments. The first is a previously-declared
slice or, as seen here, a new slice. The second is the length of the slice. The
third is the capacity of the slice, which means the length of the underlying
array. If no capacity is specified, the behavior of `make` is identical to the
above syntax. */

var yetAnotherSlice []int = make([]int, 5, 10)

/* The only two built-in functions for slices are `copy` and `append`. `copy`
copies elements from a source to a destination slice, allocating a new array if
necessary. `append` adds an arbitrary items to the end of the slice. If the
underlying array is too small to accommodate the items, a new array is created
with more space, the old array is copied into it, and a new slice is returned
that references that array. This behavior is extremely similar to arrays in
languages like JavaScript. */

/* For `copy`, the first argument is the source slice along with the items to copy.  the source and destination slices can be the same slice, different slices, or a new slice. */

func copyToSlices() {

}

/* Below, because yetAnotherSlice has nothing in it but has a specified capacity
of 10, append does not create a new array, meaning that the operation is fast.
Because anotherSlice was assigned an underlying array that was the length of the
slice, append needs to create a new, larger array before returning the slice
with 2 added, meaning that the operation is slower. */

func appendToSlices() {
	yetAnotherSlice = append(yetAnotherSlice, 4)
	anotherSlice = append(anotherSlice, 2)
}

// but with
// them all of the common array procedures from other languages can be achieved. */
