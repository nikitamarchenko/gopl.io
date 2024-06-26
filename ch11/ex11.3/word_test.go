/*

Exercise 11.3: TestRandomPalindromes only tests palindromes. Write a
randomized test that generates and verifies non-palindromes.

*/

package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23)
	n += 2
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		var r1, r2 rune
		for {
			r1 = rune(rng.Intn(0x1000))
			r2 = rune(rng.Intn(0x1000))
			if r1 == r2 {
				continue
			}
			if !unicode.IsLetter(r1) || !unicode.IsLetter(r2) {
				continue
			}
			break
		}
		runes[i] = r1
		runes[n-1-i] = r2
	}
	return string(runes)
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
