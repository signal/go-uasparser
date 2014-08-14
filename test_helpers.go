package uas

import (
  "reflect"
  "testing"
)

func Asserts(t *testing.T, what string, actual bool) {
  if !actual {
    t.Error("Expected", what, "to be", true, "but was", actual, "instead")
  }
}

func AssertEquals(t *testing.T, what string, expected interface{}, actual interface{}) {
  if expected != actual {
    t.Error("Expected", what, "to be", expected, "but was", actual, "instead")
  }
}

func AssertNil(t *testing.T, what string, actual interface{}) {
  if actual != nil {
    t.Error("Expected", what, "to be", nil, "but was", actual, "instead")
  }
}

// Deep equality checking using reflection. If the two provided types are not equal, call t.Error.
func AssertDeepEquals(t *testing.T, what string, expected interface{}, actual interface{}) {
  if !reflect.DeepEqual(expected, actual) {
    t.Error("For", what, "expected", expected, "to equal", actual)
  }
}
