package systems

import (
	"azul3d.org/native/cp.v1"
	"github.com/20zinnm/ecs"
	"github.com/20zinnm/smasaio-server/components"
)

type PhysicsSystem struct {
	gameworld *ecs.World
	world     cp.Space
	entities  []ecs.Entity
}

func (s *PhysicsSystem) Update(dt int) {

}

func (s *PhysicsSystem) ComponentMask() uint64 {
	return uint64(components.Physics)
}

func (s *PhysicsSystem) Add(entity ecs.Entity) {

}