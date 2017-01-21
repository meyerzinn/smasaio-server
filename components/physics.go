package components

import (
	"azul3d.org/native/cp.v1"
	"github.com/wooga/go-entitas"
)

type PhysicsComponent struct {
	Body *cp.Body
}

func (PhysicsComponent) Type() entitas.ComponentType {
	return Physics
}

//func (pc PhysicsData) MarshalBinary() (data []byte, err error) {
//	buffer := bytes.NewBuffer(data)
//	binary.Write(buffer, binary.LittleEndian, pc.Body.Position())
//	binary.Write(buffer, binary.LittleEndian, pc.Body.Angle())
//
//	data = buffer.Bytes()
//	return
//}
