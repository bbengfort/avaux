/**
 * Package sqrt is a trivial example package to get us started on Go
 * programming as well as BDD tests using ginkgo.
 */

package cipher

import (
	"errors"
)

// Sqrt returns an approximation of the square root of x.
func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("cipher: square root of negative number")
	}

	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}
