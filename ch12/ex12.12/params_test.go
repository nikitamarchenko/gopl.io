/*

Exercise 12.12: Extend the field tag notation to express parameter validity
requirements. For example, a string might need to be a valid email address or
credit-card number, and an integer might need to be a valid US ZIP code. Modify
Unpack to check these requirements.

*/

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"net/http"
	"strings"
	"testing"
)

func TestUnpack(t *testing.T) {

	type S struct {
		Email string `validator:"email"`
		Ccn   string `validator:"ccn"`
		Zip   int    `validator:"zip"`
	}

	var s S

	req, _ := http.NewRequest("POST", "",
		strings.NewReader("email=nm@nm.com&zip=12345&ccn=0123456789012345"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err := Unpack(req, &s); err != nil {
		t.Fatalf("Unpack() = %v", err)
	}

	t.Logf("%v", s)

	req, _ = http.NewRequest("POST", "",
		strings.NewReader("email=nonemail"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err := Unpack(req, &s); err == nil {
		t.Fatalf("Unpack() = %v", err)
	}

	req, _ = http.NewRequest("POST", "",
		strings.NewReader("ccn=invalidccn"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err := Unpack(req, &s); err == nil {
		t.Fatalf("Unpack() = %v", err)
	}

	req, _ = http.NewRequest("POST", "",
		strings.NewReader("zip=097123097"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err := Unpack(req, &s); err == nil {
		t.Fatalf("Unpack() = %v", err)
	}
}
