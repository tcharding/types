// Vector of strings
package vector

import "bytes"

type Vector []string

// Equal: true if v is equal to q
func (v Vector) Equal(q []string) bool {
	var p []string = v

	if len(p) != len(q) {
		return false
	}

	for i, _ := range p {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}

// do we need to check errors after each write to buf?
func (v Vector) String() (string, error) {
	var data []string = v
	var buf bytes.Buffer
	lastIndex := len(data) - 1

	buf.WriteRune('[')
	for i, s := range data {
		buf.WriteRune('"')
		buf.WriteString(s)
		buf.WriteRune('"')
		if i != lastIndex {
			buf.WriteString(", ")
		}
	}
	buf.WriteRune(']')

	return buf.String(), nil
}
