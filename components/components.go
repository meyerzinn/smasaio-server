package components

import (
	"github.com/20zinnm/ecs"
)

const (
	Physics ecs.Component = 1 << iota
	Input
	Ship
	Net
)
