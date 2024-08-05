package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
)

// LjposEncoder is the structure containing the writer used when encoding Envelopes
type LjposEncoder struct {
	writer io.Writer
}

// NewEncoder encode Ljpos-li packets
func NewEncoder(writer io.Writer) LjposEncoder {
	return LjposEncoder{writer: writer}
}

// Encode an Envelope packet and write the result into the writer
func (e *LjposEncoder) Encode(envelope Envelope) error {
	var payload bytes.Buffer
	encoder := gob.NewEncoder(&payload)
	if err := encoder.Encode(envelope.Payload); err != nil {
		return err
	}

	if err := binary.Write(e.writer, binary.LittleEndian, envelope.Type); err != nil {
		return err
	}
	if envelope.Size == 0 {
		envelope.Size = int32(payload.Len())
	}
	if err := binary.Write(e.writer, binary.LittleEndian, envelope.Size); err != nil {
		return err
	}
	_, err := e.writer.Write(payload.Bytes())
	if err != nil {
		return err
	}
	return nil
}
