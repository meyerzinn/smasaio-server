package world

import (
	"azul3d.org/native/cp.v1"
	"github.com/20zinnm/smasaio-server/components"
	"github.com/wooga/go-entitas"
)

type GameWorld struct {
	entitas.Pool
	*cp.Space
}

func New() *GameWorld {
	return &GameWorld{
		&entitas.NewPool(components.TOTAL, 0),
		cp.SpaceNew(),
	}
}
