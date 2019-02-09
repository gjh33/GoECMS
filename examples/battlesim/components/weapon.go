package components

// WeaponComponent is used to identify that an entity has a weapon
type WeaponComponent interface {
	Weapon() *Weapon
}

// Weapon represents how much damage an entity can do in battle
type Weapon struct {
	Name   string
	Range  float32
	Damage float32
}

// Weapon implements WeaponComponent
func (weapon *Weapon) Weapon() *Weapon {
	return weapon
}
