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
	// String conversions.
)

/*----------------------------------------------
 * Variables, pointers, and functions
 *----------------------------------------------
 */

/* Variables in Go are conceptually similar to variables in C, to wit they are
memory allocations. By default, variable declaration and use relies on
references by name, meaning that the run-time can put the value attached to that
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
or `-`. */

/* Create a pointer via the `*` prefix on the value type. The value contained at
this memory address is `nil` by default. */

var pointerToTheAnswer *int

/* A pointer to an existing value can be created with the address operator `&`.
The address operator is a unary operator that returns the memory address of
a previously declared variable.

Since the setting of a value at an address is assignment, it cannot happen
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

// Function parameters and returns can be explicitly typed.

func whoYouGonnaCall(option int) string {
	var hero string

	if option == 1 {
		hero = "Ghostbusters"
	} else if option == 2 {
		hero = "Ninja Turtles"
	} else if option == 3 {
		hero = "Jem"
	} else {
		hero = "No hero selected"
	}

	return hero
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
