package components

import (
	"time"
	"github.com/wooga/go-entitas"
)

type RegenerationComponent struct {
	Rate  float32       // How much health to regenerate per tick.
	Delay time.Duration // How long after last taking damage does the entity begin to regenerate.
}

func (RegenerationComponent) Type() entitas.ComponentType {
	return Regeneration
}
