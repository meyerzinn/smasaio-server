package systems

// The order in which systems should execute.
const (
	InputsPriority     = int(iota)
	PhysicsPriority
	NetworkingPriority
)
