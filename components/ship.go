package components

import (
	"bytes"
	"encoding/binary"
	"time"
)

const DelayBetweenShots = 3 * time.Second

type ShipComponent struct {
	NextShot     time.Time
	ShieldMillis uint32
	ShipDamage   float32
	Diamonds     uint16 // The amount of currency (diamonds) a ship has.
}

func (sc ShipComponent) MarshalBinary() (data []byte, err error) {
	var buffer bytes.Buffer
	var bools byte = 0
	if time.Now().After(sc.NextShot) {
		bools &= 1 << 1
	}
	buffer.WriteByte(bools)
	binary.Write(buffer, binary.LittleEndian, sc.ShieldMillis)

}
