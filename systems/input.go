package systems

import "time"

type InputSystem struct {
	world    *ecs.World
	entities []ecs.Entity
}

func (*InputSystem) ComponentMask() uint64 {
	return uint64(components.Input & components.Physics)
}

func (ic *InputSystem) Add(entity ecs.Entity) {
	ic.entities = append(ic.entities, entity)
}

func (ic *InputSystem) Initialize(w *ecs.World) {
	ic.world = w
}

func (ic *InputSystem) Update(dt time.Duration) {

}
