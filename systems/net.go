package systems

import (
	"github.com/20zinnm/ecs"
	"github.com/20zinnm/smasaio-server/components"
)

type NetworkingSystem struct {
	world    *ecs.World
	entities []ecs.Entity
}

func NewSystem()

func (ns *NetworkingSystem) ComponentMask() uint64 {
	return uint64(components.Player)
}

func (ns *NetworkingSystem) Add(entity ecs.Entity) {
	ns.entities = append(ns.entities, entity)
}

func (ns *NetworkingSystem) Priority() int {
	return NetworkingPriority
}

func (ns *NetworkingSystem) Update(dt int) {
	// Create world snapshots and mail them--priority mail. I have some Forever stamps around here.
	for _, entity := range ns.entities {
		ec := ns.world.EntityContainer(entity)
		if ec == nil {
			continue
		}
		//conn := ec.Components[components.Net].(components.NetData).Connection
	}
}
