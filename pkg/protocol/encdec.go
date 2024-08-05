package protocol

import "io"

type LjposEncoderDecoder struct {
	LjposDecoder
	LjposEncoder
}

func NewEncoderDecoder(rw io.ReadWriter) LjposEncoderDecoder {
	return LjposEncoderDecoder{
		NewDecoder(rw),
		NewEncoder(rw),
	}
}
