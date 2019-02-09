package components

// WeaponComponent is used to identify that an entity has a weapon
type WeaponComponent interface {
	WeaponComponent() *Weapon
}

// Weapon represents how much damage an entity can do in battle
type Weapon struct {
	Name   string
	Range  float32
	Damage float32
}

// WeaponComponent implements WeaponComponent
func (weapon *Weapon) WeaponComponent() *Weapon {
	return weapon
}
