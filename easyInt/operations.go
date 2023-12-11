// Package easyInt provides methods of operations between integers
package easyInt

// Gcd returns greatest common division of x and y
func Gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

// Floor returns the value rounded down of x divide by y
func Floor(x, y int) int {
	return x / y
}

// Ceil returns the value rounded up of x divide by y
func Ceil(x, y int) int {
	return (x + y - 1) / y
}

// Round returns the rounded value of x divide by y
func Round(x, y int) int {
	if x%y >= y/2 {
		return Ceil(x, y)
	}
	return Floor(x, y)
}
