package composition

/*
	===========================================================================================
	Proposed Solution
	===========================================================================================
	Procedure:
		- Define a function type for both the square and inc function "Operation".
		- Implement the functionality of the square and inc functions.
		- Define the compose function that implements the required function "x ↦ f(g(x))"
			= Function compose has 2 functions of type Operation as inputs and returns func with
			  the required composition as it takes an int as an input and returns the composition
			  of square(inc(num)).
		- How it will work:
			= Initialize h with compose(square, inc)
			= Now h is a function that takes int as input and returns int "func(num int) int"
			= Running h(6) leads to this structure >> "square(inc(6))"
			= inc(6) will run first >> answer will be 7
			= square(7) will run after >> final answer will be 49
*/

// Define a function type
type Operation func(int) int

// Returns square of num
func Square(num int) int {
	return num * num
}

// Increments num by 1
func Inc(num int) int {
	return num + 1
}

// compose has 2 func inputs and returns the composition func "f(g(x))"
func Compose(f Operation, g Operation) func(int) int {
	return func(num int) int {
		return f(g(num))
	}
}
