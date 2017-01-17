package systems

import (
	"github.com/20zinnm/ecs"
	"github.com/20zinnm/smasaio-server/components"
)

type PhysicsSystem struct {
	world    *ecs.World
	entities []ecs.Entity
}

func (s *PhysicsSystem) Update(dt int) {

}

func (s *PhysicsSystem) ComponentMask() uint64 {
	return uint64(components.Physics)
}

func (s *PhysicsSystem) Add(entity ecs.Entity) {

}
