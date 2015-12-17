package fields

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		age      uint16
		siblings uint16
		name     string
		phone    string
		expected []byte
	}{
		{21, 3, "Jane", "", []byte{0, 21, 0, 3, 5, 'J', 'a', 'n', 'e', 1}},
		{22, 0, "Zaphod", "5555551212", []byte{0, 22, 0, 0, 7, 'Z', 'a', 'p', 'h', 'o', 'd', 11, '5', '5', '5', '5', '5', '5', '1', '2', '1', '2'}},
	}

	for _, test := range tests {
		e := NewPersonV3Enc()
		e.SetAge(test.age)
		e.SetSiblings(test.siblings)
		e.SetName(test.name)
		e.SetPhone(test.phone)
		b := e.Bytes()
		if string(b) != string(test.expected) {
			t.Errorf("got %v want %v", b, test.expected)
		}
	}
}
