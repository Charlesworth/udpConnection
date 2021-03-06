package connectionManager

import (
	"testing"
)

func TestConnectionManager_Contains(t *testing.T) {
	testCM := NewConnectionManager()
	if testCM.ContainsIP("newIP") {
		t.Error("connection manager returned true when it did not contain that entry")
	}

	testCM.AddIP("newIP")
	if !testCM.ContainsIP("newIP") {
		t.Error("connection manager returned false when it did contain that entry")
	}
}

func TestConnectionManager_Add(t *testing.T) {
	testCM := NewConnectionManager()

	testCM.AddIP("newIP")
	if !testCM.ContainsIP("newIP") {
		t.Error("connection manager does not add unique IPs")
	}

	testCM.AddIP("newIP")
	if !testCM.ContainsIP("newIP") {
		t.Error("connection manager removed a pre existing IP on add")
	}
}

func TestConnectionManager_Remove(t *testing.T) {
	testCM := NewConnectionManager()

	testCM.AddIP("newIP")
	testCM.RemoveIP("newIP")
	if testCM.ContainsIP("newIP") {
		t.Error("connection manager did not remove IP")
	}

	//test on an empty connection manager
	testCM.RemoveIP("newIP")
	if testCM.ContainsIP("newIP") {
		t.Error("connection manager did not remove IP")
	}
}
