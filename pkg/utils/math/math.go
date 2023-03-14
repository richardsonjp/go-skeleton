package math

import (
	"fmt"
	"go-skeleton/pkg/utils/null"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Random(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// Convert string to uint.
// @TODO Bener gak nih disini?
func ConvertStringToUint(s string) (uint, error) {
	temp, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(temp), nil
}

func IsDecimalLengthValid(num *float64, maxDecimal *int) bool {
	if null.IsNil(num) || null.IsNil(maxDecimal) {
		return true
	}
	str := fmt.Sprint(*num)
	dotIndex := strings.IndexByte(str, '.')
	decimalCount := len(str) - dotIndex - 1
	return dotIndex < 0 || decimalCount <= *maxDecimal
}

// @see https://gist.github.com/cevaris/bc331cbe970b03816c6b
// @see https://floating-point-gui.de/errors/comparison/
func CompFloat64NearlyEqual(a float64, b float64) bool {
	epsilon := 0.00000001
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)

	if a == b { // shortcut, handles infinities
		return true
	} else if a == 0 || b == 0 || (absA+absB < math.SmallestNonzeroFloat32) {
		// a or b is zero or both are extremely close to it
		// relative error is less meaningful here
		return diff < (epsilon * math.SmallestNonzeroFloat32)
	} else {
		// use relative error
		return diff/math.Min((absA+absB), math.MaxFloat64) < epsilon
	}
}
