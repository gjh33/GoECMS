package ecms

var curUID = 0

// Entity interface defines what makes an object an Entity
type Entity interface {
	// Returns the unique identifier for this  entity
	UID() int
	Enabled() bool
	SetEnabled(bool)
}

// BasicEntity implements the Entity interface at a basic level
type BasicEntity struct {
	uid     int
	enabled bool
}

// NewBasicEntity is default constructor for a basic entity
func NewBasicEntity() (obj *BasicEntity) {
	obj = new(BasicEntity)
	obj.uid = curUID
	curUID++
	obj.enabled = true
	return
}

// UID from Entity interface
func (entity *BasicEntity) UID() int {
	return entity.uid
}

// Enabled from Entity interface
func (entity *BasicEntity) Enabled() bool {
	return entity.enabled
}

// SetEnabled from Entity interface
func (entity *BasicEntity) SetEnabled(val bool) {
	entity.enabled = val
}
