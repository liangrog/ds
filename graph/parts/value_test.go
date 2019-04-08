package parts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	value := "test"
	v := NewValue(value)
	assert.NotEmpty(t, v.Id())
	vs, _ := v.Value().(string)
	assert.Equal(t, vs, value)

	id := "1234"
	v.SetId(id)
	assert.Equal(t, v.Id(), id)
}
