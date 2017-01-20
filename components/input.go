package components

type InputData struct {
	// Combat
	Shooting  bool
	Shielding bool

	// Movement
	Thrusting    bool
	TurningLeft  bool
	TurningRight bool
}

//func (ic *InputData) MarshalBinary() (data []byte, err error) {
//	var encoded byte = 0
//	if ic.Shooting {
//		encoded |= 1 << 1
//	}
//	if ic.Shielding {
//		encoded |= 1 << 2
//	}
//	if ic.Thrusting {
//		encoded |= 1 << 3
//	}
//	if ic.TurningLeft {
//		encoded |= 1 << 4
//	}
//	if ic.TurningRight {
//		encoded |= 1 << 5
//	}
//	data[0] = encoded
//	return
//}
//
//func (ic *InputData) UnmarshalBinary(data []byte) error {
//	var encoded = data[0]
//	if encoded&(1<<1) > 0 {
//		ic.Shooting = true
//	}
//	if encoded&(1<<2) > 0 {
//		ic.Shielding = true
//	}
//	if encoded&(1<<3) > 0 {
//		ic.Thrusting = true
//	}
//	if encoded&(1<<4) > 0 {
//		ic.TurningLeft = true
//	}
//	if encoded&(1<<5) > 0 {
//		ic.TurningRight = true
//	}
//	return nil
//}
