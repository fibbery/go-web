package token

import (
	"testing"
)

func TestParse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1Njk0ODE1MTIsImlkIjowLCJuYmYiOjE1Njk0ODE1MTIsInVzZXJuYW1lIjoiYWRtaW4ifQ.an2Bx8Ib3bxEBG-VVxswWk7LFOmJm72M8NArzzMDCIg"
	secrect := "zfJRQ3bsC89QpEh1QxnGHxIOHZsPgU3xIbvkF2kn8y4L0RXgf6jFS5tQmRRYt0Z"
	_, e := Parse(token, secrect)
	if e != nil {
		t.Error("parse error", e)
	}

}
