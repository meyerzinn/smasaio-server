package components

import (
	"time"

	"github.com/wooga/go-entitas"
)

type DamageEvent struct {
	source entitas.EntityID
	amount float32
	time   time.Time
}

// HealthComponent represents an entity that takes and deals damage.
type HealthComponent struct {
	HealthPoints float32
	History      []DamageEvent

	// Collision specifies how much damage this entity deals when colliding with another damageable entity.
	// This is convenient because it could be reciprocal; if a player collides with a bullet, both damage each other.
	Collision float32
}