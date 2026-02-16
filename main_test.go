package main

import "testing"

func TestObjectValueToString_Simple(t *testing.T) {
	obj := ObjectValue{
		Data: map[string]interface{}{
			"age":  30,
			"city": "London",
		},
	}

	result := obj.ToString()
	expected := "{age:30,city:London}"

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
