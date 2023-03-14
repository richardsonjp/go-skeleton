package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullAsStringEmpty(t *testing.T) {
	payloadWithValue := "hello world"
	resultWithValue := NullAsStringEmpty(&payloadWithValue)
	assert.Equal(t, "hello world", resultWithValue)

	resultWithoutValue := NullAsStringEmpty(nil)
	assert.Equal(t, "", resultWithoutValue)
}

func TestStringEmptyAsNull(t *testing.T) {
	payloadWithValue := "hello world"
	resultWithValue := StringEmptyAsNull(payloadWithValue)
	assert.Equal(t, "hello world", *resultWithValue)

	resultWithoutValue := StringEmptyAsNull("")
	assert.Nil(t, resultWithoutValue)
}

func TestIsNil(t *testing.T) {
	var testInterface interface{}
	var testStringWithPointer *string
	assert.True(t, IsNil(testInterface))
	assert.True(t, IsNil(testStringWithPointer))
	assert.True(t, IsNil(nil))

	var testPointer = new(int)
	testStruct := struct {
		Test string
	}{Test: "hello"}
	assert.False(t, IsNil(testPointer))
	assert.False(t, IsNil(*testPointer))
	assert.False(t, IsNil(testStruct))
	assert.False(t, IsNil(&testStruct))
	assert.False(t, IsNil("data"))
	assert.False(t, IsNil(""))
}
