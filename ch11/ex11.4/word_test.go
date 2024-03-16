/*

Exercise 11.4: Modify randomPalindrome to exercise IsPalindrome’s handling of
punctuation and spaces.

*/

package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

var punctuation []rune

func init() {
	for r := rune(33); r < 126; r++ {
		if unicode.IsPunct(r) {
			punctuation = append(punctuation, r)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	if n == 0 {
		return ""
	}

	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		if unicode.IsPunct(r) {
			i--
			continue
		}
		runes[i] = r
		runes[n-1-i] = r
	}

	p := rng.Intn(23) + 1
	tl := p + n
	m := make([]bool, tl)
	p_count := 0
	for p_count < p {
		pos := rng.Intn(tl - 1)
		if !m[pos] {
			m[pos] = true
			p_count++
		}
	}

	res := make([]rune, tl)
	pos := 0
	for i, v := range m {
		switch v {
		case true:
			res[i] = punctuation[rng.Intn(len(punctuation)-1)]
		case false:
			res[i] = runes[pos]
			pos++
		}
	}
	return string(res)
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
