package components

import (
	"bytes"
	"encoding/binary"

	"azul3d.org/native/cp.v1"
)

type PhysicsComponent struct {
	Body cp.Body
}

func (pc PhysicsComponent) MarshalBinary() (data []byte, err error) {
	buffer := bytes.NewBuffer(data)
	binary.Write(buffer, binary.LittleEndian, pc.Body.Position())
	binary.Write(buffer, binary.LittleEndian, pc.Body.Angle())

	data = buffer.Bytes()
	return
}
