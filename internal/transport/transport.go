package transport

import (
	"encoding/binary"
	"io"
)

// for sending messages

func SendFrame(w io.Writer, data []byte) error {
	length := uint16(len(data))
	lenBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(lenBuf, length)
	_, err := w.Write(lenBuf)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// for receiving messages

func ReceiveFrame(r io.Reader) ([]byte, error) {
	var lenBuf [2]byte
	_, err := io.ReadFull(r, lenBuf[:])
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint16(lenBuf[:])
	encrypted := make([]byte, length)
	_, err = io.ReadFull(r, encrypted)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}
