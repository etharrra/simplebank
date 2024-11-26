package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcefghijklmnopqrstuvwxyz"

var randomSource *rand.Rand

/*
*
Initializes the random source using the current time as a seed value.
*/
func init() {
	randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))
}

/**
 * RandomInt generates a random integer between the specified minimum and maximum values (inclusive).
 *
 * @param min The minimum value for the random integer generation.
 * @param max The maximum value for the random integer generation.
 * @return int64 A random integer within the specified range.
 */
func RandomInt(min, max int64) int64 {
	return min + randomSource.Int63n(max-min+1)
}

/**
 * RandomString generates a random string of length n using characters from the alphabet.
 *
 * @param n The length of the random string to generate
 * @return The randomly generated string
 */
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[randomSource.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

/**
 * RandomOwner generates a random string of length 6.
 *
 * @return string: A randomly generated string of length 6.
 */
func RandomOwner() string {
	return RandomString(6)
}

/**
 * RandomMoney generates a random integer representing a monetary value between 0 and 1000.
 *
 * @return int64 - A randomly generated monetary value.
 */
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

/**
* RandomCurrency generates a random currency code from a predefined list of currencies.
*
* @return string: A randomly selected currency code from the list of currencies.
 */
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)
	return currencies[randomSource.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
