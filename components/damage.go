package components

// HealthData represents an entity that takes and deals damage.
type HealthData struct {
	HealthPoints    uint16
	MaxHealthPoints uint16
	// Collision specifies how much damage this entity deals when colliding with another damageable entity.
	// This is convenient because it could be reciprocal; if a player collides with a bullet, both damage each other.
	Collision float32
}
