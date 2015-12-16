package fields

import "fmt"

// notes/questions:
// handle endianness of uints?
// does len need to checked for size?

type PersonV2Enc struct {
	data []byte
	curFieldIndex int
}

// should this accept a value to presize the cap of data?
func NewPersonV2Enc() *PersonV2Enc {
	// Cap: should there be a default initial size?  If so, what is
	// a reasonable first approximation?  Should it be possible for
	// the caller to pass an initial cap?
	return &PersonV2Enc{data: make([]byte, 0, 32)}
}

func (e *PersonV2Enc) SetAge(v uint16) {
	if e.curFieldIndex > 0 {
		panic(fmt.Sprintf("SetAge out of order; current field index is %d", e.curFieldIndex))
	}
	e.data = append(e.data, byte(v >> 8))
	e.data = append(e.data, byte(v))
	e.curFieldIndex ++
}

func (e *PersonV2Enc) SetSiblings(v uint16) {
	if e.curFieldIndex != 1 {
		panic(fmt.Sprintf("SetSiblings out of order; current field index is %d", e.curFieldIndex))
	}
	e.data = append(e.data, byte(v >> 8))
	e.data = append(e.data, byte(v))
	e.curFieldIndex ++
}

func (e *PersonV2Enc) SetName(v string) {
	if e.curFieldIndex != 2 {
		panic(fmt.Sprintf("SetName out of order; current field index is %d", e.curFieldIndex))
	}
	e.data = append(e.data, byte(len(v) + 1))
	e.data = append(e.data, []byte(v)...)
	e.curFieldIndex ++
}

func (e *PersonV2Enc) SetPhone(v string) {
	if e.curFieldIndex != 3 {
		panic(fmt.Sprintf("SetPhone out of order; current field index is %d", e.curFieldIndex))
	}
	e.data = append(e.data, byte(len(v) + 1))
	e.data = append(e.data, []byte(v)...)
	e.curFieldIndex ++
}

func (e *PersonV2Enc) Bytes() []byte {
	return e.data
}
