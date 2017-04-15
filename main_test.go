package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetIntValue2Values(t *testing.T) {
	arr := []byte{0x02, 0x97}
	value := getIntValue(arr)
	assert.Equal(t, 663, value)
}

func TestGetIntValue3Values(t *testing.T) {
	arr := []byte{0x00, 0x0A, 0x6F}
	value := getIntValue(arr)
	assert.Equal(t, 2671, value)
}

func TestGetByteArrayWithValue(t *testing.T) {
	expected := []byte{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	value := getByteArrayWithValue(10, 5)
	assert.Equal(t, expected, value)
}
