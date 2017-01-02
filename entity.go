package manta

import (
	"fmt"

	"github.com/dotabuff/manta/dota"
)

// EntityOp is a bitmask representing the type of operation performed on an Entity
type EntityOp int

const (
	EntityOpNone           EntityOp = 0x00
	EntityOpCreated        EntityOp = 0x01
	EntityOpUpdated        EntityOp = 0x02
	EntityOpDeleted        EntityOp = 0x04
	EntityOpEntered        EntityOp = 0x08
	EntityOpLeft           EntityOp = 0x10
	EntityOpCreatedEntered EntityOp = EntityOpCreated | EntityOpEntered
	EntityOpUpdatedEntered EntityOp = EntityOpUpdated | EntityOpEntered
	EntityOpDeletedLeft    EntityOp = EntityOpDeleted | EntityOpLeft
)

var entityOpNames = map[EntityOp]string{
	EntityOpNone:           "None",
	EntityOpCreated:        "Created",
	EntityOpUpdated:        "Updated",
	EntityOpDeleted:        "Deleted",
	EntityOpEntered:        "Entered",
	EntityOpLeft:           "Left",
	EntityOpCreatedEntered: "Created+Entered",
	EntityOpUpdatedEntered: "Updated+Entered",
	EntityOpDeletedLeft:    "Deleted+Left",
}

// String returns a human identifiable string for the EntityOp
func (o EntityOp) String() string {
	return entityOpNames[o]
}

// EntityHandler is a function that receives Entity updates
type EntityHandler func(*Entity, EntityOp) error

// Entity represents a single game entity in the replay
type Entity struct {
	index   int32
	serial  int32
	class   *class
	active  bool
	state   *fieldState
	fpCache map[string]*fieldPath
}

// newEntity returns a new entity for the given index, serial and class
func newEntity(index, serial int32, class *class) *Entity {
	return &Entity{
		index:   index,
		serial:  serial,
		class:   class,
		active:  true,
		state:   newFieldState(),
		fpCache: make(map[string]*fieldPath),
	}
}

// String returns a human identifiable string for the Entity
func (e *Entity) String() string {
	return fmt.Sprintf("%d <%s>", e.index, e.class.name)
}

// Map returns a map of current entity state as key-value pairs
func (e *Entity) Map() map[string]interface{} {
	values := make(map[string]interface{})
	for _, fp := range e.class.getFieldPaths(newFieldPath(), e.state) {
		values[e.class.getNameForFieldPath(fp)] = e.state.get(fp)
	}
	return values
}

// Dump prints the current entity state to standard output
func (e *Entity) Dump() {
	_dump(e.String(), e.Map())
}

// Get returns the current value of the Entity state for the given key
func (e *Entity) Get(name string) interface{} {
	if fp, ok := e.fpCache[name]; ok {
		return e.state.get(fp)
	}

	fp := newFieldPath()
	if !e.class.getFieldPathForName(fp, name) {
		fp.release()
		return nil
	}
	e.fpCache[name] = fp

	return e.state.get(fp)
}

// GetInt32 gets given key as an int32
func (e *Entity) GetInt32(name string) (int32, bool) {
	x, ok := e.Get(name).(int32)
	return x, ok
}

// GetUint32 gets given key as a uint32
func (e *Entity) GetUint32(name string) (uint32, bool) {
	x, ok := e.Get(name).(uint32)
	return x, ok
}

// GetUint64 gets given key as a uint64
func (e *Entity) GetUint64(name string) (uint64, bool) {
	x, ok := e.Get(name).(uint64)
	return x, ok
}

// GetFloat32 gets given key as an float32
func (e *Entity) GetFloat32(name string) (float32, bool) {
	x, ok := e.Get(name).(float32)
	return x, ok
}

// GetString gets given key as a string
func (e *Entity) GetString(name string) (string, bool) {
	x, ok := e.Get(name).(string)
	return x, ok
}

// GetBool gets given key as a bool
func (e *Entity) GetBool(name string) (bool, bool) {
	x, ok := e.Get(name).(bool)
	return x, ok
}

// GetClassId returns the id of the class associated with this Entity
func (e *Entity) GetClassId() int32 {
	return e.class.classId
}

// GetClassName returns the name of the class associated with this Entity
func (e *Entity) GetClassName() string {
	return e.class.name
}

// GetIndex returns the index of this Entity
func (e *Entity) GetIndex() int32 {
	return e.index
}

func (p *Parser) onCSVCMsg_PacketEntitiesNew(m *dota.CSVCMsg_PacketEntities) error {
	r := newReader(m.GetEntityData())

	var index = int32(-1)
	var updates = int(m.GetUpdatedEntries())
	var cmd uint32
	var classId int32
	var serial int32
	var e *Entity
	var op EntityOp

	if !m.GetIsDelta() {
		if p.newEntityFullPackets > 0 {
			return nil
		}
		p.newEntityFullPackets++
	}

	for ; updates > 0; updates-- {
		index += int32(r.readUBitVar()) + 1
		op = EntityOpNone

		cmd = r.readBits(2)
		if cmd&0x01 == 0 {
			if cmd&0x02 != 0 {
				classId = int32(r.readBits(p.classIdSize))
				serial = int32(r.readBits(17))
				r.readVarUint32()

				class := p.newClassesById[classId]
				if class == nil {
					_panicf("unable to find new class %d", classId)
				}

				baseline := p.newClassBaselines[classId]
				if baseline == nil {
					_panicf("unable to find new baseline %d", classId)
				}

				e = newEntity(index, serial, class)
				p.newEntities[index] = e
				readFields(newReader(baseline), class.serializer, e.state)
				readFields(r, class.serializer, e.state)
				op = EntityOpCreated | EntityOpEntered

			} else {
				if e = p.newEntities[index]; e == nil {
					_panicf("unable to find existing entity %d", index)
				}

				op = EntityOpUpdated
				if !e.active {
					e.active = true
					op |= EntityOpEntered
				}

				readFields(r, e.class.serializer, e.state)
			}

		} else {
			if e = p.newEntities[index]; e == nil {
				_panicf("unable to find existing entity %d", index)
			}

			if !e.active {
				_panicf("entity %d (%s) ordered to leave, already inactive", e.class.classId, e.class.name)
			}

			op = EntityOpLeft
			if cmd&0x02 != 0 {
				op |= EntityOpDeleted
				p.newEntities[index] = nil
			}
		}

		for _, h := range p.newEntityHandlers {
			if err := h(e, op); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *Parser) OnEntity(h EntityHandler) {
	p.newEntityHandlers = append(p.newEntityHandlers, h)
}
