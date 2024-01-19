package main

import "testing"

func TestCase1(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Peek()
	param_3 := obj.Pop()
	param_4 := obj.Empty()
	if param_2 != 1 {
		t.Errorf("Expected %v, got %v", 1, param_2)
	}
	if param_3 != 1 {
		t.Errorf("Expected %v, got %v", 1, param_3)
	}
	if param_4 != false {
		t.Errorf("Expected %v, got %v", false, param_4)
	}
}

func TestCase2(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	param_1 := obj.Pop()
	param_2 := obj.Empty()
	if param_1 != 1 {
		t.Errorf("Expected %v, got %v", 1, param_1)
	}
	if param_2 != true {
		t.Errorf("Expected %v, got %v", true, param_2)
		t.Log(obj)
	}
}
