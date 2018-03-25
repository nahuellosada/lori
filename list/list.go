// Package list provides functions that operates on slices .
// Each function name follows the format : FunctionnameType where type is the type of slice contents
package list

// I is short for integer
type I func(int) int

// S is short for string
type S func(string) string

// F32 is short for float
type F32 func(float32) float32

// AdjustInt applies a function to the item at Index of a slice
func AdjustInt(f I, index int, slice []int) []int {

	r := make([]int, len(slice))
	r = slice[:]
	tmp := f(r[index])
	r[index] = tmp

	return r

}

// AdjustStr applies a function (f) to the item at Index of a slice
func AdjustStr(f S, index int, slice []string) []string {
	r := make([]string, len(slice))
	r = slice[:]
	tmp := f(r[index])
	r[index] = tmp

	return r
}

// AdjustF32 applies a function to the item at Index of a slice
func AdjustF32(f F32, index int, slice []float32) []float32 {

	r := make([]float32, len(slice))
	r = slice[:]
	tmp := f(r[index])
	r[index] = tmp

	return r

}
