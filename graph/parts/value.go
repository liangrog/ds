package parts

import (
	"github.com/google/uuid"
	"github.com/liangrog/ds/graph/intrfs"
)

var _ intrfs.Value = &Value{}

type Value struct {
	id string

	value interface{}
}

func NewValue(v interface{}) *Value {
	return &Value{
		id:    uuid.New().String(),
		value: v,
	}
}

func (v *Value) SetId(i string) *Value {
	v.id = i
	return v
}

func (v *Value) Id() string {
	return v.id
}

func (v *Value) SetValue(value interface{}) *Value {
	v.value = value
	return v
}

func (v *Value) Value() interface{} {
	return v.value
}

func (v *Value) DeepCopy() *Value {
	return &Value{
		id:    v.id,
		value: v.value,
	}
}
