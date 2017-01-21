package components

import (
	"bytes"
	"encoding/binary"
	"time"
	"github.com/wooga/go-entitas"
)

type ShipComponent struct {
	LastShot          time.Time
	DelayBetweenShots time.Duration

	ShieldActive    bool
	ShieldRemaining time.Duration
}

func (ShipComponent) Type() entitas.ComponentType {
	return Ship
}

//func (sc ShipData) MarshalBinary() (data []byte, err error) {
//	var buffer bytes.Buffer
//	var bools byte = 0
//	if time.Now().After(sc.NextShot) {
//		bools &= 1 << 1
//	}
//	buffer.WriteByte(bools)
//	binary.Write(buffer, binary.LittleEndian, sc.ShieldMillis)
//
//}
