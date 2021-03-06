package main

import "testing"

func TestAdd(context *testing.T) {
	var first int = 13
	var second int = 3

	var result int = Add(first, second)

	if result != 16 {
		context.Errorf("Expected 16 got %d", result)
	} else {
		context.Log("13 + 6 = 16")
	}

	first = 5
	second = 8

	result = Add(first, second)

	if result != 13 {
		context.Errorf("Expected 13 got %d", result)
	} else {
		context.Log("5 + 8 = 13")
	}
}
