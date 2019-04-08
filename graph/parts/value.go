package parts

import (
	"github.com/google/uuid"
	"github.com/liangrog/ds/graph/intrfs"
)

// Comply to Value interface
var _ intrfs.Value = &Value{}

// Value struct.
type Value struct {
	// Internal value Id.
	id string

	// Any values
	value interface{}
}

// Value constructor. Initiated with a build in ID.
func NewValue(v interface{}) *Value {
	return &Value{
		id:    uuid.New().String(),
		value: v,
	}
}

// Function for override the internal ID.
func (v *Value) SetId(i string) *Value {
	v.id = i
	return v
}

// Get internal ID.
func (v *Value) Id() string {
	return v.id
}

// Set value to struct.
func (v *Value) SetValue(value interface{}) *Value {
	v.value = value
	return v
}

// Returns the value from struct.
func (v *Value) Value() interface{} {
	return v.value
}

// Deep copy the value struct. Maintain the same internal ID.
func (v *Value) DeepCopy() *Value {
	return &Value{
		id:    v.id,
		value: v.value,
	}
}
