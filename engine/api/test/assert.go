package test

import (
	"fmt"
	"testing"

	"github.com/fsamin/go-dump"
	"github.com/ovh/cds/sdk"
	"github.com/stretchr/testify/assert"
)

// NoError logs Fatal if there is an error.
func NoError(t *testing.T, err error, msg ...interface{}) {
	if !assert.NoError(t, err) {
		t.Fatal(msg...)
	}
}

// Error logs Fatal if there is no error.
func Error(t *testing.T, err error, msg ...interface{}) {
	if !assert.Error(t, err) {
		t.Fatal(msg...)
	}
}

// NotNil logs Fatal if there is nil value.
func NotNil(t *testing.T, i interface{}, msg ...interface{}) {
	if !assert.NotNil(t, i) {
		t.Fatal(msg...)
	}
}

// Nil logs Fatal if there is no nil value.
func Nil(t *testing.T, i interface{}, msg ...interface{}) {
	if !assert.Nil(t, i) {
		t.Fatal(msg...)
	}
}

// NotEmpty logs Fatal if it's empty
func NotEmpty(t *testing.T, i interface{}, msg ...interface{}) {
	if !assert.NotEmpty(t, i) {
		t.Fatal(msg...)
	}
}

// ArrayContains check if an element exists in an array using DeepEquals function
func ArrayContains(array interface{}, s interface{}) bool {
	b := sdk.InterfaceSlice(array)
	for _, i := range b {
		if DeepEquals(i, s) {
			return true
		}
	}
	return false
}

// Equal checks that two elements are equal.
func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if !DeepEquals(expected, actual) {
		t.Log("Expected:" + dump.MustSdump(expected))
		t.Log("Actual:" + dump.MustSdump(actual))
		assert.FailNow(t, "Equal failed", msgAndArgs...)
	}
}

// NotEqual checks that two elements are not equal.
func NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if DeepEquals(expected, actual) {
		t.Log("Expected:" + dump.MustSdump(expected))
		t.Log("Actual:" + dump.MustSdump(actual))
		assert.FailNow(t, "Not equal failed", msgAndArgs...)
	}
}

// DeepEquals returns equality between 2 elements using github.com/fsamin/go-dump
func DeepEquals(a, b interface{}) bool {
	m1, err := dump.ToMap(a)
	if err != nil {
		return false
	}
	m2, err := dump.ToMap(b)
	if err != nil {
		return false
	}

	for k, v := range m1 {
		if v2, ok := m2[k]; !ok {
			fmt.Printf("%s not found\n", k)
			return false
		} else if v != v2 {
			return false
		}
	}

	for k, v := range m2 {
		if v2, ok := m1[k]; !ok {
			return false
		} else if v != v2 {
			return false
		}
	}

	return true
}

// EqualValuesWithoutOrder checks equality between two slices without respecting slide order
func EqualValuesWithoutOrder(t *testing.T, a, b interface{}, msgAndArgs ...interface{}) {
	s1 := sdk.InterfaceSlice(a)
	s2 := sdk.InterfaceSlice(b)

	for _, x := range s1 {
		if !ArrayContains(s2, x) {
			assert.Fail(t, "EqualValuesWithoutOrder failed", msgAndArgs...)
		}
	}

	for _, x := range s2 {
		if !ArrayContains(s1, x) {
			assert.Fail(t, "EqualValuesWithoutOrder failed", msgAndArgs...)
		}
	}

	if t.Failed() {
		t.Logf("Expected: %v", a)
		t.Logf("Actual: %v", b)
		assert.FailNow(t, "Equal failed", msgAndArgs...)
	}
}
