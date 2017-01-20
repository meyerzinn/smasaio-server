package systems

// The order in which systems should execute.
const (
	PhysicsPriority    = int(iota)
	NetworkingPriority
)
