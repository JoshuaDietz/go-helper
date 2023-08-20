package gohelpers

// MapFunc is a function that takes a slice of type A and applies the given function to each element
// Then it returns a slice of type B
func MapFunc[A any, B any](arr []A, f func(A) B) []B {
	result := make([]B, 0, len(arr))
	for _, v := range arr {
		result = append(result, f(v))
	}
	return result
}
