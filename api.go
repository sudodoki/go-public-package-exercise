package gopublicpackageexercise

import "golang.org/x/exp/constraints"

// Add function takes two input arguments and returns their sum
//
// It has two parameters: a and b, both ints
// No error is returned
// More info on addition at [mathisfun]
//
// [mathisfun]: https://www.mathsisfun.com/numbers/addition.html

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](a, b T) T {
	return a + b
}
